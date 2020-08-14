package main

import(
	"net/http"
	"log"
	"fmt"
	"strings"
	"html/template"
)
type myMux struct{

}
func sayhello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintln(w, "hello my route")
}

func login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	}else{
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])

	}
}
func (mux *myMux)ServeHTTP(w http.ResponseWriter, r *http.Request){
		   if r.URL.Path == "/hello"{
			   sayhello(w, r)
			   return 
		   } else if r.URL.Path =="/login"{
			   login(w, r)
			   return 
		   }
		   
}
func main(){
	//http.HandleFunc("/hello", sayhello)
	//http.HandleFunc('')
	mux := &myMux{}
	err := http.ListenAndServe(":8080", mux)
	if err  != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}