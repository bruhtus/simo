package utils_test

import (
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
						"got %v, want %v",
						output, tt.output,
					)
				}
			},
		)
	}
}
