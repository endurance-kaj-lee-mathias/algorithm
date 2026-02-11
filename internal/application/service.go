package application

import (
	"context"
	"strconv"
	"strings"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/domain"
)

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ComputeStress(ctx context.Context, m domain.HealthMetric) (int, error) {
	score := 50

	if m.HeartRate > 0 {
		if m.HeartRate > 100 {
			delta := (m.HeartRate - 100) / 2
			if delta > 40 {
				delta = 40
			}
			score += delta
		} else if m.HeartRate < 60 {
			delta := (60 - m.HeartRate) / 2
			if delta > 20 {
				delta = 20
			}
			score -= delta
		}
	}

	if m.SleepTimeMinutes > 0 {
		if m.SleepTimeMinutes < 420 {
			delta := (420 - m.SleepTimeMinutes) / 14
			if delta > 30 {
				delta = 30
			}
			score += delta
		} else if m.SleepTimeMinutes > 540 {
			delta := (m.SleepTimeMinutes - 540) / 30
			if delta > 10 {
				delta = 10
			}
			score -= delta
		}
	}

	if bp := strings.TrimSpace(m.BloodPressure); bp != "" {
		if before, after, ok := strings.Cut(bp, "/"); ok {
			sysStr := strings.TrimSpace(before)
			diaStr := strings.TrimSpace(after)
			sys, err1 := strconv.Atoi(sysStr)
			dia, err2 := strconv.Atoi(diaStr)
			if err1 == nil && err2 == nil {
				if sys >= 160 || dia >= 100 {
					score += 40
				} else if sys >= 140 || dia >= 90 {
					score += 30
				} else if sys >= 130 || dia >= 85 {
					score += 20
				} else if sys < 110 && dia < 70 {
					score -= 5
				}
			}
		}
	}

	if score < 0 {
		score = 0
	} else if score > 100 {
		score = 100
	}
	return score, nil
}
