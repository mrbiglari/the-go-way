package main

import "fmt"

func nonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case message := <-messages:
		fmt.Println("message received:", message)
	default:
		fmt.Println("no message received")
	}

	message := "message"
	select {
	case messages <- message: // message cannot be sent to the messages channel, because the channel has no buffer and there is no receiver. Therefore the default case is selected.
		fmt.Println("sent message", message)
	default:
		fmt.Println("no message sent")
	}

	select {
	case message := <-messages: // checks if messages has data to receive, it doesn’t because the previous send never happened.
		fmt.Println("message received", message)
	case <-signals:
		fmt.Println("signal received")
	default:
		fmt.Println("no activity")
	}
}
