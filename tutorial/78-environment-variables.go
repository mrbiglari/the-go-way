package main

import (
	"fmt"
	"os"
	"strings"
)

func environmentVariables() {
	print := fmt.Println

	os.Setenv("API_KEY", "bob_went_to_sydney")
	print(os.Getenv("API_KEY"))
	print(os.Getenv("DATABASE_SETTINGS"))

	for _, variable := range os.Environ() {
		print(variable)
		keyValuePair := strings.SplitN(variable, "=", 2)
		print(keyValuePair[0], ":", keyValuePair[1])
	}
}
