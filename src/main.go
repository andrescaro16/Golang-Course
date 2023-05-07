package main
import (
	"fmt"
	"sync"
)

func print(text string, wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println(text)
}

func counter(wg *sync.WaitGroup)  {
	defer wg.Done()
	for i:= 0; i<10; i++ {
		fmt.Println(i)
	}
}

func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go print("Hola", &wg)
	go counter(&wg)
	wg.Wait()
}