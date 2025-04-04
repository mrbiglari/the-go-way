package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func logging() {
	log.Println("standard logging")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	customLogger := log.New(os.Stdout, "custom prefix:", log.LstdFlags)
	customLogger.Println("bob went to sydney")

	customLogger.SetPrefix("new custom prefix:")
	customLogger.Println("sydney went to bob creek")

	var buffer bytes.Buffer
	bufferLogger := log.New(&buffer, "buffered log:", log.LstdFlags)
	bufferLogger.Println("bob went to sydney again!")

	fmt.Println(buffer.String())

	handler := slog.NewJSONHandler(os.Stderr, nil)
	structuredLogger := slog.New(handler)
	structuredLogger.Info("bob and sydney went to sydney")
	structuredLogger.Info("json log", "key", "val", "age", 42)
}
