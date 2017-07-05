package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    "strings"
)

func sendWithEndpointMethodAsGet() {
	res, err := http.Post("http://127.0.0.1/ingest.php","application/json",strings.NewReader("{\"endpoint\":{ \"method\":\"GET\",\"url\":\"http://127.0.0.1/thirdPartyReceiveFromDelivery.php/data?title={mascot}&image={location}&foo={bar}\" },\"data\":[ {\"mascot\":\"Gopher\",\"location\":\"https://blog.golang.org/gopher/gopher.png\" }] }"))
        if err != nil {
                log.Fatal(err)
        }
        responseBody, err := ioutil.ReadAll(res.Body)
        res.Body.Close()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s", responseBody)
}

func sendWithEndpointMethodAsPost() {
	res, err := http.Post("http://127.0.0.1/ingest.php","application/json",strings.NewReader("{\"endpoint\":{ \"method\":\"POST\",\"url\":\"http://127.0.0.1/thirdPartyReceiveFromDelivery.php/\" },\"data\":[ {\"mascot\":\"Gopher\",\"location\":\"https://blog.golang.org/gopher/gopher.png\" }] }"))
        if err != nil {
                log.Fatal(err)
        }
        responseBody, err := ioutil.ReadAll(res.Body)
        res.Body.Close()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s", responseBody)
}

func main() {
	sendWithEndpointMethodAsGet()
	sendWithEndpointMethodAsPost()

}
