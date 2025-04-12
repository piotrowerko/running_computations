package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/piotrowerko/running_computations/keycomputations"
	"github.com/piotrowerko/running_computations/pkg/constants"
)


func main() {
	distanceFlag := flag.Float64("distance", 0.0, "Distance in kilometers")
	timeFlag := flag.Int("time", 0, "Time in seconds")
	timeFormatFlag := flag.String("timeformat", "", "Time in HH:MM:SS format")
	intervalFlag := flag.Float64("interval", 1.0, "Interval in kilometers for timestamps")
	presetFlag := flag.String("preset", "", "Preset distance (5k, 10k, half, marathon)")

	flag.Parse()

	distance := *distanceFlag

	// Sprawdzenie czy podano preset
	if *presetFlag != "" {
		presetLower := strings.ToLower(*presetFlag)
		switch presetLower {
		case "5k":
			distance = constants.FiveK
		case "10k":
			distance = constants.TenK
		case "half":
			distance = constants.HalfMarathon
		case "marathon":
			distance = constants.Marathon
		default:
			fmt.Printf("Nieznany preset: %s. Dostępne presety: 5k, 10k, half, marathon\n", *presetFlag)
			os.Exit(1)
		}
		fmt.Printf("Użyto predefiniowanego dystansu: %.3f km\n", distance)
	}

	// Obsługa czasu podanego w różnych formatach
	timeInSeconds := *timeFlag
	if *timeFormatFlag != "" {
		var err error
		timeInSeconds, err = keycomputations.ParseTimeFormat(*timeFormatFlag)
		if err != nil {
			fmt.Printf("Błąd: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Podany czas: %s (%d sekund)\n", *timeFormatFlag, timeInSeconds)
	}

	if distance <= 0 || timeInSeconds <= 0 {
		fmt.Println("Proszę podać prawidłowe wartości dystansu i czasu")
		fmt.Println("Użycie: cli -distance=10.0 -time=3600")
		fmt.Println("lub: cli -distance=10.0 -timeformat=01:00:00")
		fmt.Println("lub: cli -preset=10k -timeformat=01:00:00")
		os.Exit(1)
	}

	pace := keycomputations.ComputePace(distance, timeInSeconds)
	minutes := int(pace)
	seconds := int((pace - float64(minutes)) * 60)
	fmt.Printf("Tempo biegu: %d min %d sek na kilometr\n", minutes, seconds)

	// Obliczanie i wyświetlanie timestampów
	timestamps := keycomputations.ComputeTimeStamps(distance, timeInSeconds, *intervalFlag, keycomputations.ComputePace)
	timestampsStr := keycomputations.ConvertTimeStamps(timestamps)

	fmt.Println("\nCzasy na poszczególnych odcinkach:")
	for i, ts := range timestampsStr {
		fmt.Printf("%.1f km: %s\n", *intervalFlag*float64(i+1), ts)
	}
}
