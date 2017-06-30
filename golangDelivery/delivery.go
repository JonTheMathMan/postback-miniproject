package main

import (
	"fmt"
	"strings"
	"github.com/garyburd/redigo/redis"
	
	"io/ioutil"
	"log"
	"net/http"
	"time"

)

func getRequestMaps() (map[string]string, map[string]string) {
	// connect to redis with redigo
	redisConn,err := redis.Dial("tcp","127.0.0.1:8037")
	redisConn.Do("AUTH","IlikeredFROGS865IhateredFROGS865")
	
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

func placeRequestMapsIntoResponse()(string,string) {
    endpoint,data := getRequestMaps()
    
    responseBody := endpoint["url"]
    retrievedKeys := getUrlDataKeys(responseBody)
    for i:=0;i<len(retrievedKeys);i++ {
        replaceByUrlKey := "{"+retrievedKeys[i]+"}"
        responseBody = strings.Replace(responseBody,replaceByUrlKey,data[retrievedKeys[i]],1)
    }
    
    return endpoint["method"],responseBody
}

func sendAGetResponse() {
    method,body := placeRequestMapsIntoResponse()
        
    if(body!="" && method=="GET") {
        thirdPartyResponse, err := http.Get(body)
        if err != nil {
            log.Fatal(err)
        }
        thirdPartyResponseBody, err := ioutil.ReadAll(thirdPartyResponse.Body)
        thirdPartyResponse.Body.Close()
        if err != nil {
            log.Fatal(err)
        }

        //log.Print(thirdPartyResponseBody)
	fmt.Printf("%s", thirdPartyResponseBody)
    }
}

func continuousFuncCall(repeatFunc func()) {
	time.Sleep(100 * time.Millisecond)
	repeatFunc()
	continuousFuncCall(repeatFunc)
}

func main() {
        continuousFuncCall(sendAGetResponse)
}
