package utils_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bruhtus/simo/utils"
)

func TestIsExpired(t *testing.T) {
	now := time.Now().Round(time.Second)

	endTimeCases := []struct {
		endTime time.Time
		output  bool
	}{
		{now.Add(time.Duration(-1 * time.Second)), true},
		{now.Add(time.Duration(-1 * time.Minute)), true},
		{now.Add(time.Duration(-1 * time.Hour)), true},
		{now.Add(time.Duration(0 * time.Second)), true},
		{now.Add(time.Duration(1 * time.Second)), false},
		{now.Add(time.Duration(1 * time.Minute)), false},
		{now.Add(time.Duration(1 * time.Hour)), false},
	}

	for _, tt := range endTimeCases {
		t.Run(
			tt.endTime.Format("2006-01-02_15:04:05"),
			func(t *testing.T) {
				output := utils.DetermineIsExpired(tt.endTime)
				if output != tt.output {
					t.Errorf(
						"got %t, want %t",
						output, tt.output,
					)
				}
			},
		)
	}
}

func TestGetDurationMinutesAndSeconds(t *testing.T) {
	durationCases := []struct {
		duration string
		output   string
	}{
		{"2h15m20s", "135:20"},
		{"69m42s", "69:42"},
		{"60m69s", "61:09"},
	}

	for _, tt := range durationCases {
		t.Run(
			tt.duration,
			func(t *testing.T) {
				parsedDuration, err := time.ParseDuration(tt.duration)
				if err != nil {
					t.Fatalf("Error parsing duration: %v", err)
				}

				minutes, seconds := utils.GetDurationMinutesAndSeconds(
					parsedDuration,
				)

				output := fmt.Sprintf("%02d:%02d", minutes, seconds)

				if output != tt.output {
					t.Errorf(
						"got %s, want %s",
						output, tt.output,
					)
				}
			},
		)
	}
}

func TestRemainingDuration(t *testing.T) {
	now := time.Now().Round(time.Second)

	endTimeCases := []struct {
		endTime time.Time
		output  string
	}{
		{now.Add(time.Duration(1 * time.Second)), "00:01"},
		{now.Add(time.Duration(1 * time.Minute)), "01:00"},
		{now.Add(time.Duration(1 * time.Hour)), "60:00"},
	}

	for _, tt := range endTimeCases {
		t.Run(
			tt.endTime.Format("2006-01-02_15:04:05"),
			func(t *testing.T) {
				minutes, seconds := utils.ParseRemainingDuration(
					tt.endTime,
				)

				output := fmt.Sprintf("%02d:%02d", minutes, seconds)

				if output != tt.output {
					t.Errorf(
						"got %s, want %s",
						output, tt.output,
					)
				}
			},
		)
	}
}
