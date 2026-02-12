package models

import (
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/domain"
)

type MessageModel struct {
	HeartRate        int    `json:"heartRate"`
	SleepTimeMinutes int    `json:"sleepTimeMinutes"`
	BloodPressure    string `json:"bloodPressure"`
}

func ToModel(msg domain.HealthMetric) MessageModel {
	return MessageModel{
		HeartRate:        msg.HeartRate,
		SleepTimeMinutes: msg.SleepTimeMinutes,
		BloodPressure:    msg.BloodPressure,
	}
}

type StressModel struct {
	Stress int `json:"stress"`
}
