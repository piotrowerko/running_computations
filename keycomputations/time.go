package keycomputations

import (
	"fmt"
	"strings"
	"time"

	"github.com/piotrowerko/running_computations/pkg/constants"
)

// ParseTimeFormat converts time format HH:MM:SS to seconds
func ParseTimeFormat(timeFormat string) (int, error) {
	parts := strings.Split(timeFormat, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid time format, use HH:MM:SS")
	}

	t, err := time.Parse("15:04:05", timeFormat)
	if err != nil {
		return 0, fmt.Errorf("cannot parse time: %v", err)
	}

	hours := t.Hour()
	minutes := t.Minute()
	seconds := t.Second()

	totalSeconds := hours*constants.SecondsInHour + minutes*constants.SecondsInMinute + seconds
	return totalSeconds, nil
}
