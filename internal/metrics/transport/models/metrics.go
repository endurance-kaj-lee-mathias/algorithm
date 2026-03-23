package models

import "time"

type SampleModel struct {
	HeartRate        float64   `json:"heartRate"`
	RMSSD            float64   `json:"rmssd"`
	RestingHeartRate float64   `json:"restingHeartRate"`
	Steps            float64   `json:"steps,omitempty"`
	SleepDebtHours   float64   `json:"sleepDebtHours,omitempty"`
	RecordedAt       time.Time `json:"recordedAt"`
}

type ComputeRequest struct {
	UserID  string        `json:"userId"`
	Samples []SampleModel `json:"samples"`
}

type ComputeResponse struct {
	Score    float64 `json:"score"`
	Category string  `json:"category"`
	ZHR      float64 `json:"zHr"`
	ZRMSSD   float64 `json:"zRmssd"`
}
