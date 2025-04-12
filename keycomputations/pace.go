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

// funkcja ComputeTimeStamps przyjmuje dystans i czas i odcinek i funkcję do obliczenia tempa (ComputePace) co ile km ma byc timestamp, zwraca slice z timestampami
func ComputeTimeStamps(distance float64, timeInSeconds int, distInterval float64, paceFunc func(float64, int) float64) []int {
	timeStamps := []int{}

	pace := paceFunc(distance, timeInSeconds)

	// Liczba interwałów
	numIntervals := int(distance / distInterval)

	// in loop add new timestamps:= i*dictInterval * pace * constants.SecondsInMinute
	for i := 1; i <= numIntervals; i++ {
		timeStamps = append(timeStamps, int(float64(i)*distInterval*pace*constants.SecondsInMinute))
	}

	return timeStamps
}

// funkcja ConvertTimeStamps przyjmuje slice z timestampami w sekundach i zwraca slice z timestampami w formacie HH:MM:SS
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
