package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// messages := make(chan string, 1)
	// signals := make(chan bool, 1)

	select {
	case msg := <-messages:
		fmt.Println("received message ", msg)
	default:
		fmt.Println("no received message")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Printf("message: (%s) sent", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message ", msg)
	case sig := <-signals:
		fmt.Printf("signal: (%+v) received", sig)
	default:
		fmt.Println("no activity")
	}

}
