package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encoding() {
	_string := "abc123!?$*&()'-=@~"
	base64EncodedString := base64.StdEncoding.EncodeToString([]byte(_string))
	fmt.Println(base64EncodedString)

	decodedString, _ := base64.StdEncoding.DecodeString(base64EncodedString)
	fmt.Println(string(decodedString))

	urlEncodedString := base64.URLEncoding.EncodeToString([]byte(_string))
	fmt.Println(urlEncodedString)

	decodedUrlString, _ := base64.URLEncoding.DecodeString(urlEncodedString)
	fmt.Println(string(decodedUrlString))
}
