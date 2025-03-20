package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello!")
}

func headers(writer http.ResponseWriter, request *http.Request) {

	for name, headers := range request.Header {
		for _, header := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, header)
		}
	}
}

func httpServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	_error := http.ListenAndServe(":8090", nil)
	if _error != nil {
		panic(_error)
	}
}
