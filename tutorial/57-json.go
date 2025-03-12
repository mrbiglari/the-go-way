package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var fruits = []string{"apple", "peach", "pear"}
var fruitStock = map[string]int{"apple": 5, "pear": 10}
var resultsJson = `{"page":1, "fruits":["berry", "orange", "banana"]}`

func _json() {
	//Marshal
	boolJson, _ := json.Marshal(true)
	fmt.Println(string(boolJson))

	intJson, _ := json.Marshal(12)
	fmt.Println(string(intJson))

	floatJson, _ := json.Marshal(42.42)
	fmt.Println(string(floatJson))

	stringJson, _ := json.Marshal("gopher")
	fmt.Println(string(stringJson))

	stringArrayJson, _ := json.Marshal(fruits)
	fmt.Println(string(stringArrayJson))

	mapJson, _ := json.Marshal(fruitStock)
	fmt.Println(string(mapJson))

	_response1 := &response1{Page: 1, Fruits: fruits}
	response1Json, _ := json.Marshal(_response1)
	fmt.Println(string(response1Json))

	_response2 := &response2{Page: 1, Fruits: fruits}
	response2Json, _ := json.Marshal(_response2)
	fmt.Println(string(response2Json))

	//Unmarshal
	jsonData := []byte(`{"number":6.9, "strings":["bob", "sydney"]}`)
	var mapData map[string]interface{}
	_error := json.Unmarshal(jsonData, &mapData)
	if _error != nil {
		panic(_error)
	}
	fmt.Println(mapData)
	number, _ := mapData["number"].(float64)
	fmt.Println(number)

	_strings, _ := mapData["strings"].([]interface{})
	_string, _ := _strings[0].(string)
	fmt.Println(_string)

	unmarshaledResponse := response2{}
	json.Unmarshal([]byte(resultsJson), &unmarshaledResponse)
	fmt.Println(unmarshaledResponse)
	fmt.Println(unmarshaledResponse.Fruits[0])

	//Encode
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(fruitStock)

	//Decode
	decoder := json.NewDecoder(strings.NewReader(resultsJson))
	response := response2{}
	decoder.Decode(&response)
	fmt.Println(response)
}
