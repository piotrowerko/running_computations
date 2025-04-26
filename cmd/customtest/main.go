package main

import (
	"github.com/piotrowerko/running_computations/keycomputations"
)

func main() {
	// Wywołaj funkcję PrintSum z pakietu keycomputations
	keycomputations.PrintSum()

	// Użyj eksportowanych zmiennych i funkcji
	result := keycomputations.SumNumbers(keycomputations.Number1, keycomputations.Number2)
	println("Obliczona suma:", result)
}
