package main

import (
    "fmt"
    "net/http"
)

type myMux struct{

}

func (mux *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/" {
		sayHello(w,r)
		return 
	}

	http.NotFound(w, r)
	return 
}

func sayHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "wujinbo yes!")
	
}

func main(){
	mux := &myMux{}
	http.ListenAndServe(":8080", mux)
}