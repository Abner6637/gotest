package main

import (
	"fmt"
	"gotest/calc"
)

func main() {
	fmt.Println("Hello, World!")
	result := calc.Add(3, 5)
	result2 := calc.Add(2, 4)
	fmt.Println("Result:", result)
	fmt.Println("Result:", result2)
}
