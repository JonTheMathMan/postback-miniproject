package main

import (
	"fmt"
	"strings"
	"github.com/garyburd/redigo/redis"
	
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
	"os"

)

func dialRedis() redis.Conn {
	// connect to redis with redigo
        redisConn,err := redis.Dial("tcp","127.0.0.1:8037")
	if err!=nil {
		fmt.Errorf("problem dialing redis")
	}
        if _,errAuth := redisConn.Do("AUTH","IlikeredFROGS865IhateredFROGS865"); errAuth!=nil {
		fmt.Errorf("problem with redis authorization")
	}
	return redisConn
}

func getRequestMaps(redisConn redis.Conn) (map[string]string, map[string]string) {

    //get the look up id's
    requestLookUpId,err := redis.String(redisConn.Do("RPOP","queueList"))
    requestEndpointLookUpId,err := redis.String(redisConn.Do("HGET",requestLookUpId,"endpoint"))
    requestDataLookUpId,err := redis.String(redisConn.Do("HGET",requestLookUpId,"data"))
    
    //make the request maps
    endpoint,err := redis.StringMap(redisConn.Do("HGETALL",requestEndpointLookUpId))
    data,err := redis.StringMap(redisConn.Do("HGETALL",requestDataLookUpId))
    
    //clean up the redis database
    redisConn.Do("DEL",requestLookUpId)
    redisConn.Do("DEL",requestEndpointLookUpId)
    redisConn.Do("DEL",requestDataLookUpId)
    
    if(err!=nil) {
        fmt.Print("err is not nil")
    }
    
    return endpoint,data
}

func getUrlDataKeys(s string) []string {
	read:=false
	firstKey:=true
	
	reading := func(r rune) rune {
		switch {
			case r=='{' && firstKey==false:
				read=true
				return ','
			case r=='{':
				read = true
				return -1
			case r=='}':
				read=false
				firstKey=false
				return -1
			case read==true:
				return r
			default:
				return -1
		}
	}
	
	keysString := strings.Map(reading,s)
	return strings.Split(keysString, ",")
}

func placeDataMapIntoGetResponse(urlWithKeys string, data map[string]string) string {
    retrievedKeys := getUrlDataKeys(urlWithKeys)
    for i:=0;i<len(retrievedKeys);i++ {
        replaceByUrlKey := "{"+retrievedKeys[i]+"}"
        urlWithKeys = strings.Replace(urlWithKeys,replaceByUrlKey,data[retrievedKeys[i]],1)
    }
    
    return urlWithKeys
}

func placeDataIntoTypeUrlValues(data map[string]string) url.Values {
	dataValues := url.Values{}
	for key,value := range data {
		dataValues.Add(key,value)
	}
	return dataValues
}

func sendAResponse(redisConn redis.Conn) {
	
	file, err := os.OpenFile("deliveryLog.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()
    
    deliveryLogger := log.New(file, "deliveryLogger: ", log.Llongfile)
	
	endpoint,data := getRequestMaps(redisConn)
        
    if(endpoint["method"]=="GET") {
	body := placeDataMapIntoGetResponse(endpoint["url"],data)
	sendTime := time.Now()
        thirdPartyResponse, err := http.Get(body)
        if err != nil {
            deliveryLogger.Fatal(err)
        }
        thirdPartyResponseBody, err := ioutil.ReadAll(thirdPartyResponse.Body)
        thirdPartyResponseStatusCode := thirdPartyResponse.StatusCode
	thirdPartyResponse.Body.Close()
        if err != nil {
            deliveryLogger.Fatal(err)
        }
	receiveResponseTime := time.Now()



	deliveryLogger.Println(endpoint["method"])
	deliveryLogger.Println(body)
	deliveryLogger.Println("result:")
	deliveryLogger.Println("response code:",thirdPartyResponseStatusCode)
	deliveryLogger.Println("response time:",receiveResponseTime.Sub(sendTime).Nanoseconds(),"nanoseconds")
	deliveryLogger.Printf("%s", thirdPartyResponseBody)
    }else if endpoint["method"]=="POST" {
	dataAsFormValues := placeDataIntoTypeUrlValues(data)
        sendTime := time.Now()
        thirdPartyResponse, err := http.PostForm(endpoint["url"],dataAsFormValues)
        if err != nil {
            deliveryLogger.Fatal(err)
        }
        thirdPartyResponseBody, err := ioutil.ReadAll(thirdPartyResponse.Body)
        thirdPartyResponseStatusCode := thirdPartyResponse.StatusCode
        thirdPartyResponse.Body.Close()
        if err != nil {
            deliveryLogger.Fatal(err)
        }
        receiveResponseTime := time.Now()



        deliveryLogger.Println(endpoint["method"])
        deliveryLogger.Println(endpoint["url"])
	deliveryLogger.Println(dataAsFormValues)
        deliveryLogger.Println("result:")
        deliveryLogger.Println("response code:",thirdPartyResponseStatusCode)
        deliveryLogger.Println("response time:",receiveResponseTime.Sub(sendTime).Nanoseconds(),"nanoseconds")
        deliveryLogger.Printf("%s", thirdPartyResponseBody)
    }
}

func continuousFuncCall(repeatFunc func(redis.Conn),redisConn redis.Conn) {
	time.Sleep(100 * time.Millisecond)
	repeatFunc(redisConn)
	continuousFuncCall(repeatFunc,redisConn)
}

func main() {
	redisConn := dialRedis()
        continuousFuncCall(sendAResponse,redisConn)
}
