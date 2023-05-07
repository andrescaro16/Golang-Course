package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Max"] = 19
	m["Andrés"] = 23

	fmt.Println(m)

	// Range the map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// To see a key's value
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}