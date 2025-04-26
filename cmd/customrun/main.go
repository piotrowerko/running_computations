package main

import "fmt"

var number1 = int64(10)
var number2 = int64(20)

func sumNumbers(a int64, b int64) int64 {
	return a + b
}

func main() {
	result := sumNumbers(number1, number2)
	fmt.Println("Suma liczb:", result)
}
