package keycomputations

import (
	"fmt"

	"github.com/piotrowerko/running_computations/pkg/constants"
)

func ComputePace(distance float64, timeInSeconds int) float64 {
	timeInMinutes := float64(timeInSeconds) / constants.SecondsInMinute
	pace := timeInMinutes / distance
	return pace
}

// ComputeTimeStamps takes distance, time, interval and pace calculation function (ComputePace) as parameters,
// and returns a slice of timestamps for each kilometer interval
func ComputeTimeStamps(distance float64, timeInSeconds int, distInterval float64, paceFunc func(float64, int) float64) []int {
	timeStamps := []int{}

	pace := paceFunc(distance, timeInSeconds)

	// Number of intervals
	numIntervals := int(distance / distInterval)

	// in loop add new timestamps:= i*dictInterval * pace * constants.SecondsInMinute
	for i := 1; i <= numIntervals; i++ {
		timeStamps = append(timeStamps, int(float64(i)*distInterval*pace*constants.SecondsInMinute))
	}

	return timeStamps
}

// ComputeTimeStampsNegativeSplit calculates timestamps for negative split strategy
// where runner runs slower in the first part of the distance and faster in the second part
func ComputeTimeStampsNegativeSplit(distance float64,
	timeInSeconds int,
	distInterval float64,
	splitDistancePercentage int16,
	paceDifferencePercentage int16,
	paceFunc func(float64, int) float64) []int {
	timeStamps := []int{}

	// Calculate average pace for the entire distance
	avgPace := paceFunc(distance, timeInSeconds)

	// Calculate the split point distance
	splitPoint := distance * float64(splitDistancePercentage) / 100.0

	// Calculate pace for slower and faster parts
	paceRatio := float64(paceDifferencePercentage) / 100.0

	// To maintain average pace across the entire distance, we need to balance paces appropriately
	// We assume: slowerPartPace = avgPace * (1 + adjustment) and fasterPartPace = avgPace * (1 - adjustment)
	// Where adjustment = paceRatio * (distance - splitPoint) / distance
	adjustment := paceRatio * splitPoint / distance
	slowerPartPace := avgPace * (1.0 + adjustment)
	fasterPartPace := avgPace * (1.0 - adjustment*splitPoint/(distance-splitPoint))

	// Number of intervals
	numIntervals := int(distance / distInterval)

	// Variable to keep track of current time
	currentTime := 0.0

	for i := 1; i <= numIntervals; i++ {
		// Calculate distance at the end of this interval
		intervalEnd := float64(i) * distInterval

		// Check if interval belongs entirely to the slower part
		if intervalEnd <= splitPoint {
			// Entire interval in slower part
			currentTime += distInterval * slowerPartPace * constants.SecondsInMinute
		} else if intervalEnd > splitPoint && intervalEnd-distInterval < splitPoint {
			// Interval partially in slower part, partially in faster part
			distInSlowerPart := splitPoint - (intervalEnd - distInterval)
			distInFasterPart := distInterval - distInSlowerPart

			currentTime += distInSlowerPart * slowerPartPace * constants.SecondsInMinute
			currentTime += distInFasterPart * fasterPartPace * constants.SecondsInMinute
		} else {
			// Entire interval in faster part
			currentTime += distInterval * fasterPartPace * constants.SecondsInMinute
		}

		// Add time to the timestamps list
		timeStamps = append(timeStamps, int(currentTime))
	}

	return timeStamps
}

// ConvertTimeStamps takes a slice of timestamps in seconds and returns a slice of timestamps in HH:MM:SS format
func ConvertTimeStamps(timeStamps []int) []string {
	timeStampsStr := []string{}

	for _, timestamp := range timeStamps {
		hours := timestamp / constants.SecondsInHour
		minutes := (timestamp % constants.SecondsInHour) / constants.SecondsInMinute
		seconds := timestamp % constants.SecondsInMinute
		timeStampsStr = append(timeStampsStr, fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds))
	}

	return timeStampsStr
}
