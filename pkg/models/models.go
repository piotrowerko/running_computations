package models

type PaceInputModel struct {
	Distance       float64 `json:"distance"`
	TimeInSeconds  int     `json:"timeInSeconds,omitempty"`
	TimeFormat     string  `json:"timeFormat,omitempty"`
	Preset         string  `json:"preset,omitempty"`
	Interval       float64 `json:"interval,omitempty"`
	NegativeSplit  bool    `json:"negativeSplit,omitempty"`
	SplitDistance  int     `json:"splitDistance,omitempty"`
	PaceDifference int     `json:"paceDifference,omitempty"`
}

type PaceOutputModel struct {
	Distance            float64        `json:"distance"`
	TimeInSeconds       int            `json:"timeInSeconds"`
	PaceMinutes         int            `json:"paceMinutes"`
	PaceSeconds         int            `json:"paceSeconds"`
	PaceDecimal         float64        `json:"paceDecimal"`
	StrategyDescription string         `json:"strategyDescription"`
	Intervals           []IntervalData `json:"intervals"`
}

type IntervalData struct {
	Distance float64 `json:"distance"`
	Time     string  `json:"time"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
