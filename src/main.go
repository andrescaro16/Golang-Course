package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	// This line will be executed until the end of the function, and thus frees the goroutine of the WaitGroup
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	// The sync package allows to interact in a primitive way with the goroutines. Variable that compiles a set of gorutines and releases them little by little
	var wg sync.WaitGroup

	fmt.Println("Hello")

	// We indicate that we are going to add 1 goroutine to the WaitGroup so that it waits for its execution before the base goroutine (main) dies, and thus gives time to the next goroutine to run
	wg.Add(1)

	// The reserved word go will execute the function concurrently (Goroutine)
	go say("world", &wg)

	// WaitGroup function used to tell the main goroutine (main) to wait until all the goroutines of the WaitGroup are finished, that is, until 'defer wg.Done()' is executed in each of the goroutines
	wg.Wait()

	// Anonymous function
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	// Function so that when it reaches this line it waits the indicated time (long enough for the Gorutine to execute its function concurrently). Sleep() is not good practice
	time.Sleep(time.Second * 1)
}
