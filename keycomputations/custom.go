package keycomputations

import "fmt"

// Number1 - eksportowana zmienna
var Number1 = int64(10)

// Number2 - eksportowana zmienna
var Number2 = int64(20)

// SumNumbers - eksportowana funkcja do sumowania
func SumNumbers(a int64, b int64) int64 {
	return a + b
}

// PrintSum - funkcja drukująca sumę
func PrintSum() {
	fmt.Println(SumNumbers(Number1, Number2))
}
