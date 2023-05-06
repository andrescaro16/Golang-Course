package main

import "fmt"

func normalFunction(name string) {
	fmt.Printf("Hey %s, how are you?\n", name)
}

func tripleArgument(num, num2 int, name string) string {
	return fmt.Sprintf("Has %s %d or %d years old?", name, num, num2)
}

func doubleReturn(num int) (int, int) {
	return num, num * 2
}

func main() {
	normalFunction("Andrew")

	question := tripleArgument(1, 5, "Camilo")
	fmt.Println(question)

	_, double := doubleReturn(5)
	fmt.Println(double)
}
