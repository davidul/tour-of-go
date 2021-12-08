package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello\n")
}
