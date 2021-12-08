package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", loadBalancer)
	http.ListenAndServe(":8090", nil)
}

func loadBalancer(res http.ResponseWriter, req *http.Request)  {
	log.Println(req.Method)
	log.Println(req.URL)
	log.Println(req.Host)
	log.Println(req)
}