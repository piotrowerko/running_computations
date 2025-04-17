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
	distanceFlag := flag.Float64("distance", 0.0, "Custom distance in kilometers (e.g., 3.3)")
	timeFlag := flag.Int("time", 0, "Time in seconds")
	timeFormatFlag := flag.String("timeformat", "", "Time in HH:MM:SS format")
	intervalFlag := flag.Float64("interval", 1.0, "Interval in kilometers for timestamps")
	presetFlag := flag.String("preset", "", "Preset distance (5k, 10k, half, marathon)")
	negativeSplitFlag := flag.Bool("negativesplit", false, "Use negative split pacing strategy")
	splitDistanceFlag := flag.Int("splitdistance", 50, "Percentage of distance for split point (default 50%)")
	paceDifferenceFlag := flag.Int("pacedifference", 5, "Percentage difference between slower and faster pace (default 5%)")

	flag.Parse()

	distance := *distanceFlag

	// Check if preset is provided
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
			fmt.Printf("Unknown preset: %s\n", *presetFlag)
			fmt.Println("Available presets: 5k, 10k, half, marathon")
			fmt.Println("Or use -distance flag for custom distance (e.g., -distance=3.3)")
			os.Exit(1)
		}
		fmt.Printf("Using preset distance: %.3f km\n", distance)
	} else if distance > 0 {
		fmt.Printf("Using custom distance: %.3f km\n", distance)
	}

	// Handle time in different formats
	timeInSeconds := *timeFlag
	if *timeFormatFlag != "" {
		var err error
		timeInSeconds, err = keycomputations.ParseTimeFormat(*timeFormatFlag)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Provided time: %s (%d seconds)\n", *timeFormatFlag, timeInSeconds)
	}

	if distance <= 0 || timeInSeconds <= 0 {
		fmt.Println("Please provide valid distance and time values")
		fmt.Println("Usage examples:")
		fmt.Println("  cli -distance=3.3 -time=1200")
		fmt.Println("  cli -distance=3.3 -timeformat=00:20:00")
		fmt.Println("  cli -preset=10k -timeformat=00:45:00")
		fmt.Println("  cli -preset=10k -timeformat=00:45:00 -negativesplit -splitdistance=60 -pacedifference=3")
		os.Exit(1)
	}

	pace := keycomputations.ComputePace(distance, timeInSeconds)

	// Converting pace format from decimal value to minutes and seconds
	paceMinutes := int(pace)
	paceSeconds := int((pace - float64(paceMinutes)) * 60)
	fmt.Printf("Running pace: %d min %d sec / km\n", paceMinutes, paceSeconds)

	// Calculate and display timestamps based on selected strategy
	var timestamps []int
	var strategyDescription string

	if *negativeSplitFlag {
		timestamps = keycomputations.ComputeTimeStampsNegativeSplit(
			distance,
			timeInSeconds,
			*intervalFlag,
			int16(*splitDistanceFlag),
			int16(*paceDifferenceFlag),
			keycomputations.ComputePace,
		)
		strategyDescription = fmt.Sprintf("Negative split strategy (%.0f%% distance point, %.0f%% pace difference)",
			float64(*splitDistanceFlag), float64(*paceDifferenceFlag))
	} else {
		timestamps = keycomputations.ComputeTimeStamps(
			distance,
			timeInSeconds,
			*intervalFlag,
			keycomputations.ComputePace,
		)
		strategyDescription = "Even pace strategy"
	}

	timestampsStr := keycomputations.ConvertTimeStamps(timestamps)

	fmt.Printf("\nPacing strategy: %s\n", strategyDescription)
	fmt.Println("Times at each interval:")
	for i, ts := range timestampsStr {
		fmt.Printf("%.1f km: %s\n", *intervalFlag*float64(i+1), ts)
	}
}
