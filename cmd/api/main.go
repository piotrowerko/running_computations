package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/piotrowerko/running_computations/keycomputations"
	"github.com/piotrowerko/running_computations/pkg/constants"
	"github.com/piotrowerko/running_computations/pkg/models"
)

const port = ":8080"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := models.HealthResponse{
		Status:  "ok",
		Message: "API is working properly",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func paceHandler(w http.ResponseWriter, r *http.Request) {
	// Check if method is POST
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  "error",
			Message: "Method not allowed. Use POST.",
		})
		return
	}

	// Decode request
	var request models.PaceInputModel
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Process request
	distance := request.Distance

	// Handle preset
	if request.Preset != "" {
		switch request.Preset {
		case "5k":
			distance = constants.FiveK
		case "10k":
			distance = constants.TenK
		case "half":
			distance = constants.HalfMarathon
		case "marathon":
			distance = constants.Marathon
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Status:  "error",
				Message: "Unknown preset: " + request.Preset,
			})
			return
		}
	}

	// Handle time
	timeInSeconds := request.TimeInSeconds
	if request.TimeFormat != "" {
		var err error
		timeInSeconds, err = keycomputations.ParseTimeFormat(request.TimeFormat)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Status:  "error",
				Message: "Invalid time format: " + err.Error(),
			})
			return
		}
	}

	// Validate input
	if distance <= 0 || timeInSeconds <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status:  "error",
			Message: "Please provide valid distance and time values",
		})
		return
	}

	// Set default interval if not provided
	interval := request.Interval
	if interval <= 0 {
		interval = 1.0
	}

	// Set default split distance if not provided
	splitDistance := request.SplitDistance
	if splitDistance <= 0 {
		splitDistance = 50
	}

	// Set default pace difference if not provided
	paceDifference := request.PaceDifference
	if paceDifference <= 0 {
		paceDifference = 5
	}

	// Compute pace
	pace := keycomputations.ComputePace(distance, timeInSeconds)
	paceMinutes := int(pace)
	paceSeconds := int((pace - float64(paceMinutes)) * 60)

	// Calculate timestamps based on strategy
	var timestamps []int
	var strategyDescription string

	if request.NegativeSplit {
		timestamps = keycomputations.ComputeTimeStampsNegativeSplit(
			distance,
			timeInSeconds,
			interval,
			int16(splitDistance),
			int16(paceDifference),
			keycomputations.ComputePace,
		)
		strategyDescription = "Negative split strategy (" +
			strconv.Itoa(splitDistance) + "% distance point, " +
			strconv.Itoa(paceDifference) + "% pace difference)"
	} else {
		timestamps = keycomputations.ComputeTimeStamps(
			distance,
			timeInSeconds,
			interval,
			keycomputations.ComputePace,
		)
		strategyDescription = "Even pace strategy"
	}

	// Convert timestamps to HH:MM:SS format
	timestampsStr := keycomputations.ConvertTimeStamps(timestamps)

	// Build response with intervals
	intervals := make([]models.IntervalData, len(timestampsStr))
	for i, ts := range timestampsStr {
		intervals[i] = models.IntervalData{
			Distance: interval * float64(i+1),
			Time:     ts,
		}
	}

	// Build response
	response := models.PaceOutputModel{
		Distance:            distance,
		TimeInSeconds:       timeInSeconds,
		PaceMinutes:         paceMinutes,
		PaceSeconds:         paceSeconds,
		PaceDecimal:         pace,
		StrategyDescription: strategyDescription,
		Intervals:           intervals,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/pace", paceHandler)

	log.Printf("Starting API server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
