package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type film struct {
	Title string `json:"title"`
}

func httpClient() {
	check := func(_error error) {
		if _error != nil {
			panic(_error)
		}
	}
	response, _error1 := http.Get("https://gobyexample.com")
	check(_error1)

	defer response.Body.Close()
	fmt.Println("Response status: ", response.Status)

	apiResponse, _error2 := http.Get("https://swapi.dev/api/films/1/")
	check(_error2)
	defer apiResponse.Body.Close()

	byteResponse, _ := io.ReadAll(apiResponse.Body)
	var _film film
	_error3 := json.Unmarshal(byteResponse, &_film)
	check(_error3)

	fmt.Println(_film)
}
