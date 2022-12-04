package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			writer.WriteHeader(200)
			writer.Write([]byte("Hello"))
		}
	})

	http.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("This is foo"))
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func raw() {
	//server := http.Server{
	//	Addr:              ":9090",
	//	Handler:           nil,
	//	TLSConfig:         nil,
	//	ReadTimeout:       0,
	//	ReadHeaderTimeout: 0,
	//	WriteTimeout:      0,
	//	IdleTimeout:       0,
	//	MaxHeaderBytes:    0,
	//	TLSNextProto:      nil,
	//	ConnState:         nil,
	//	ErrorLog:          nil,
	//	BaseContext:       nil,
	//	ConnContext:       nil,
	//}
}
