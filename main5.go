package main

import (
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){
	requestUrl := "http://www.baidu.com"
	response, err := http.Get(requestUrl)
	if err != nil{
		fmt.Println("err:", err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(body))
    }

    fmt.Println("-----------------------------------------")
    fmt.Printf("response：%+v\n", response)
    fmt.Println("-----------------------------------------")
    fmt.Printf("response.Body：%+v\n", response.Body)
    fmt.Printf("response.Header：%+v\n", response.Header)
    fmt.Printf("response.StatusCode：%+v\n", response.StatusCode)
    fmt.Printf("response.Status：%+v\n", response.Status)
    fmt.Printf("response.Request：%+v\n", response.Request)
	fmt.Printf("response.Cookies：%+v\n", response.Cookies())
}