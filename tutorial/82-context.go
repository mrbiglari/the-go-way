package main

import (
	"fmt"
	"net/http"
	"time"
)

func bye(writer http.ResponseWriter, request *http.Request) {
	context := request.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("goodbye!")
	case <-context.Done():
		_error := context.Err()
		fmt.Println(_error) // the context’s Err() method returns an error that explains why the Done() channel was closed.
		errorStatus := http.StatusInternalServerError
		http.Error(writer, _error.Error(), errorStatus)
	}

}
func context() {
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":8090", nil)
}
