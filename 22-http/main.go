package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Sample handler function
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

	wrapper := LoggingWrapper{}
	headerWrapper := HeaderWrapper{}
	http.Handle("/handle",
		headerWrapper.Wrap(
			wrapper.Wrap(
				&MyHandler{})))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}

}

type MyHandler struct {
}

func (m MyHandler) ServeHTTP(r http.ResponseWriter, req *http.Request) {
	r.Write([]byte("Handler"))
}

type Wrapper interface {
	Wrap(h http.Handler) http.Handler
}

type LoggingWrapper struct {
}

type HeaderWrapper struct {
}

func (lw LoggingWrapper) Wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("hello")
		handler.ServeHTTP(w, req)
	})
}

func (hw HeaderWrapper) Wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("X-Custom", "Some value")
		handler.ServeHTTP(w, req)
	})
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
