package main

import (
	"crypto/sha256"
	"fmt"
)

func sha256Hashes() {
	_string := "sha256 this string please!"
	hasher := sha256.New()
	hasher.Write([]byte(_string))

	hashed := hasher.Sum(nil)

	fmt.Println(_string)
	fmt.Printf("%x\n", hashed)
}
