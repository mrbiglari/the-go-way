package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signals() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // UNIX: signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println("\nReceived signal:", sig)
		done <- true
	}()

	fmt.Println("Awaiting signal (Press Ctrl+C to exit)...")
	<-done
	fmt.Println("Exiting gracefully")
}
