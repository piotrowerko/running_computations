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

func ComputeTimeStampsNegativeSplit(distance float64,
	timeInSeconds int,
	distInterval float64,
	splitDistanceProcentage int16,
	paceDifrancePercentage int16,
	paceFunc func(float64, int) float64) []int {
	timeStamps := []int{}

	// Wyliczenie średniego tempa na cały dystans
	avgPace := paceFunc(distance, timeInSeconds)

	// Wyliczenie punktu podziału dystansu
	splitPoint := distance * float64(splitDistanceProcentage) / 100.0

	// Wyliczenie tempa dla wolniejszej i szybszej części
	paceRatio := float64(paceDifrancePercentage) / 100.0

	// Aby zachować średnie tempo na całym dystansie, musimy odpowiednio zrównoważyć tempa
	// Zakładamy: slowerPartPace = avgPace * (1 + adjustment) i fasterPartPace = avgPace * (1 - adjustment)
	// Gdzie adjustment = paceRatio * (distance - splitPoint) / distance
	adjustment := paceRatio * splitPoint / distance
	slowerPartPace := avgPace * (1.0 + adjustment)
	fasterPartPace := avgPace * (1.0 - adjustment*splitPoint/(distance-splitPoint))

	// Liczba interwałów
	numIntervals := int(distance / distInterval)

	// Zmienna przechowująca aktualny czas
	currentTime := 0.0

	for i := 1; i <= numIntervals; i++ {
		// Obliczenie dystansu na końcu tego interwału
		intervalEnd := float64(i) * distInterval

		// Sprawdzamy czy interwał należy całkowicie do wolniejszej części
		if intervalEnd <= splitPoint {
			// Cały interwał w wolniejszej części
			currentTime += distInterval * slowerPartPace * constants.SecondsInMinute
		} else if intervalEnd > splitPoint && intervalEnd-distInterval < splitPoint {
			// Interwał częściowo w wolniejszej, częściowo w szybszej części
			distInSlowerPart := splitPoint - (intervalEnd - distInterval)
			distInFasterPart := distInterval - distInSlowerPart

			currentTime += distInSlowerPart * slowerPartPace * constants.SecondsInMinute
			currentTime += distInFasterPart * fasterPartPace * constants.SecondsInMinute
		} else {
			// Cały interwał w szybszej części
			currentTime += distInterval * fasterPartPace * constants.SecondsInMinute
		}

		// Dodanie czasu do listy timestampów
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
