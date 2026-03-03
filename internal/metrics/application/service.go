package application

import (
	"context"
	"errors"
	"math"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/domain"
)

const minSamples = 12
const epsilon = 1e-9

var ErrNotEnoughSamples = errors.New("at least 12 samples are required")

func (s *service) ComputeStress(_ context.Context, _ string, samples []domain.Sample) (domain.StressResult, error) {
	if len(samples) < minSamples {
		return domain.StressResult{}, ErrNotEnoughSamples
	}

	hrs := make([]float64, len(samples))
	rmssds := make([]float64, len(samples))

	for i, sample := range samples {
		hrs[i] = sample.HeartRate
		rmssds[i] = sample.RMSSD
	}

	meanHR := mean(hrs)
	stdHR := stddev(hrs, meanHR)

	meanRMSSD := mean(rmssds)
	stdRMSSD := stddev(rmssds, meanRMSSD)

	last := len(samples) - 1

	zHR := zScore(samples[last].HeartRate, meanHR, stdHR)

	zRMSSD := -zScore(samples[last].RMSSD, meanRMSSD, stdRMSSD)

	weighted := 0.6*zHR + 0.4*zRMSSD

	sleepDebt := samples[last].SleepDebtHours
	if sleepDebt > 0 {
		weighted += 0.1 * math.Min(sleepDebt/8.0, 1.0)
	}

	score := sigmoid(weighted) * 100

	return domain.StressResult{
		Score:    math.Round(score*10) / 10,
		Category: categorize(score),
		ZHR:      math.Round(zHR*100) / 100,
		ZRMSSD:   math.Round(zRMSSD*100) / 100,
	}, nil
}

func mean(vals []float64) float64 {
	sum := 0.0
	for _, v := range vals {
		sum += v
	}
	return sum / float64(len(vals))
}

func stddev(vals []float64, m float64) float64 {
	if len(vals) < 2 {
		return 0
	}

	sum := 0.0
	for _, v := range vals {
		d := v - m
		sum += d * d
	}

	// Sample standard deviation (n-1)
	return math.Sqrt(sum / float64(len(vals)-1))
}

func zScore(val, m, std float64) float64 {
	if std < epsilon {
		return 0
	}
	return (val - m) / std
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func categorize(score float64) string {
	switch {
	case score < 35:
		return "Low"
	case score < 65:
		return "Moderate"
	default:
		return "High"
	}
}
