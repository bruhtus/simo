package scmd_test

import (
	"testing"
	"time"

	scmd "github.com/bruhtus/simo/internal"
	"github.com/bruhtus/simo/utils"
)

func TestPause(t *testing.T) {
	file := utils.TestSetupTempFile(t)

	remainingDurationCases := []struct {
		duration time.Duration
		output   string
	}{
		{time.Duration(1 * time.Second), "0m1s"},
		{time.Duration(1 * time.Minute), "1m0s"},
		{time.Duration(1 * time.Hour), "60m0s"},
	}

	for _, tt := range remainingDurationCases {
		t.Run(
			tt.duration.String(),
			func(t *testing.T) {
				status := utils.Status{
					State:      utils.StateFocus,
					IsNotify:   true,
					PausePoint: nil,
					EndTime:    time.Now().Add(tt.duration),
				}

				utils.TestSetupStatusFile(t, status, file)
				scmd.Pause(file.Name())

				resultJSON := utils.ReadStatusFile(file.Name())
				if resultJSON.PausePoint == nil {
					t.Errorf(
						"Got nil, want %s",
						tt.output,
					)
				}
			},
		)
	}

	t.Cleanup(func() {
		err := file.Close()
		if err != nil {
			t.Fatalf(
				"Failed to close file: %v",
				err,
			)
		}
	})
}
