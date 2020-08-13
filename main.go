package main

import(
	"net/http"
	"log"
	"fmt"
	"strings"
)

/**
他需要2个参数，一个是http.ResponseWriter，另一个是http.Request
往http.ResponseWriter写入什么内容，浏览器的网页源码就是什么内容。
http.Request里面是封装了，浏览器发过来的请求（包含路径、浏览器类型等等）。
 */

 func sayHello(w http.ResponseWriter, r *http.Request){
	 fmt.Println("----------------")
	 r.ParseForm()
	 fmt.Println(r.Form)
	 fmt.Println("path: ", r.URL.Path)
	 fmt.Println("scheme: ", r.URL.Scheme)
	 fmt.Println(r.Form["url_long"])
	 for k,v := range r.Form{
		 fmt.Println("key: ", k)
		 fmt.Println("val: ", strings.Join(v, ""))
	 }
	 fmt.Fprintf(w, "hello GoWeb!")
 }

 func main(){
	 http.HandleFunc("/wujinbo", sayHello)
	 err := http.ListenAndServe(":8080", nil)
	 if err != nil{
		 log.Fatal("ListenAndServe: ", err)
	 }
 }