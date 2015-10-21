package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)


type ReqJson struct {
	Name string
}

type PostResponse struct {
	Greeting string
}
 


func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var reqJson ReqJson
	err := decoder.Decode(&reqJson)
	if err != nil {
		panic("Error decoding json.")
	}
	postResp := PostResponse{Greeting: "Hello, " + reqJson.Name + "!"}
	json.NewEncoder(rw).Encode(postResp)
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello/", helloPost)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}