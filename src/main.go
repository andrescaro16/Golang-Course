package main

import (
	"fmt"
)

func interaction(channel chan string)  {
	var text string
	fmt.Scanln(&text)
	channel <- text
}

func main()  {
	channel := make(chan string, 1)
	go interaction(channel)
	message := <- channel
	fmt.Println(message)
}