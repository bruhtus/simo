package scmd_test

import (
	"testing"
	"time"

	scmd "github.com/bruhtus/simo/internal"
	"github.com/bruhtus/simo/utils"
)

func TestPause(t *testing.T) {
	dirPath := t.TempDir()
	file := utils.TestSetupTempFile(t, dirPath)

	t.Cleanup(func() {
		err := file.Close()
		if err != nil {
			t.Fatalf(
				"Failed to close file: %v",
				err,
			)
		}
	})

	remainingDurationCases := []struct {
		duration                time.Duration
		remainingDurationOutput string
		isNotifyInput           bool
		isNotifyOutput          bool
	}{
		{
			time.Duration(1 * time.Second), "0m1s",
			true, true,
		},
		{
			time.Duration(1 * time.Minute), "1m0s",
			false, false,
		},
		{
			time.Duration(1 * time.Hour), "60m0s",
			true, true,
		},
	}

	for _, tt := range remainingDurationCases {
		t.Run(
			tt.duration.String(),
			func(t *testing.T) {
				status := utils.Status{
					State:      utils.StateFocus,
					IsNotify:   tt.isNotifyInput,
					PausePoint: nil,
					EndTime:    time.Now().Add(tt.duration),
				}

				utils.TestSetupStatusFile(t, status, file)
				scmd.Pause(file.Name())

				resultJSON := utils.ReadStatusFile(file.Name())
				if resultJSON.PausePoint == nil {
					t.Errorf(
						"Got nil, want %s",
						tt.remainingDurationOutput,
					)
				}

				if resultJSON.IsNotify != tt.isNotifyOutput {
					t.Errorf(
						"Got %t, want %t",
						tt.isNotifyInput,
						tt.isNotifyOutput,
					)
				}
			},
		)
	}
}
