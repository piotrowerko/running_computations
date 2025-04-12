package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/piotrowerko/running_computations/keycomputations"
)

func main() {
	distanceFlag := flag.Float64("distance", 0.0, "Distance in kilometers")
	timeFlag := flag.Int("time", 0, "Time in seconds")
	intervalFlag := flag.Float64("interval", 1.0, "Interval in kilometers for timestamps")

	flag.Parse()

	if *distanceFlag <= 0 || *timeFlag <= 0 {
		fmt.Println("Please provide valid distance and time values")
		fmt.Println("Usage: cli -distance=10.0 -time=3600")
		os.Exit(1)
	}

	pace := keycomputations.ComputePace(*distanceFlag, *timeFlag)
	fmt.Printf("Tempo biegu: %.2f minut na kilometr\n", pace)

	// Obliczanie i wyświetlanie timestampów
	timestamps := keycomputations.ComputeTimeStamps(*distanceFlag, *timeFlag, *intervalFlag, keycomputations.ComputePace)
	timestampsStr := keycomputations.ConvertTimeStamps(timestamps)

	fmt.Println("\nCzasy na poszczególnych odcinkach:")
	for i, ts := range timestampsStr {
		fmt.Printf("%.1f km: %s\n", *intervalFlag*float64(i+1), ts)
	}
}
