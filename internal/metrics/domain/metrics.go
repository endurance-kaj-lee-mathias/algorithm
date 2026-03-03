package domain

import "time"

type Sample struct {
	HeartRate        float64
	RMSSD            float64
	RestingHeartRate float64
	Steps            float64
	SleepDebtHours   float64
	RecordedAt       time.Time
}

type StressResult struct {
	Score    float64
	Category string
	ZHR      float64
	ZRMSSD   float64
}
