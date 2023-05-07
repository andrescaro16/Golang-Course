package main

import "fmt"

func say(text string, channel chan<- string) {
	channel <- text
}

func main() {
	channel := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", channel)

	fmt.Println(<-channel)
}