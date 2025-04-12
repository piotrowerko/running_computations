package keycomputations

import (
	"fmt"
	"strings"
	"time"

	"github.com/piotrowerko/running_computations/pkg/constants"
)


// parseTimeFormat przetwarza format czasu HH:MM:SS na sekundy
func ParseTimeFormat(timeFormat string) (int, error) {
	parts := strings.Split(timeFormat, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("nieprawidłowy format czasu, użyj HH:MM:SS")
	}

	t, err := time.Parse("15:04:05", timeFormat)
	if err != nil {
		return 0, fmt.Errorf("nie można przetworzyć czasu: %v", err)
	}

	hours := t.Hour()
	minutes := t.Minute()
	seconds := t.Second()

	totalSeconds := hours*constants.SecondsInHour + minutes*constants.SecondsInMinute + seconds
	return totalSeconds, nil
}