package models

import "time"

type SampleModel struct {
	HeartRate        float64   `json:"heart_rate"`
	RMSSD            float64   `json:"rmssd"`
	RestingHeartRate float64   `json:"resting_heart_rate"`
	Steps            float64   `json:"steps,omitempty"`
	SleepDebtHours   float64   `json:"sleep_debt_hours,omitempty"`
	RecordedAt       time.Time `json:"recorded_at"`
}

type ComputeRequest struct {
	UserID  string        `json:"user_id"`
	Samples []SampleModel `json:"samples"`
}

type ComputeResponse struct {
	Score    float64 `json:"score"`
	Category string  `json:"category"`
	ZHR      float64 `json:"z_hr"`
	ZRMSSD   float64 `json:"z_rmssd"`
}
