package main

import "fmt"

func main() {
	value := sum(3, 6)
	fmt.Println(value)
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func division(a int, b int) int {
	return a / b
}
