package scmd_test

import (
	"testing"
	"time"

	scmd "github.com/bruhtus/simo/internal"
	"github.com/bruhtus/simo/utils"
)

func TestStatusPause(t *testing.T) {
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

	pausePointCases := []struct {
		input  string
		output string
	}{
		{"0m1s", "PF00:01"},
		{"1m0s", "PF01:00"},
		{"60m0s", "PF60:00"},
	}

	for _, tt := range pausePointCases {
		t.Run(
			tt.input,
			func(t *testing.T) {
				remainingDuration, err := time.ParseDuration(
					tt.input,
				)

				if err != nil {
					t.Fatalf(
						"Error parsing duration: %v",
						err,
					)
				}

				status := utils.Status{
					State:      utils.StateFocus,
					IsNotify:   false,
					PausePoint: &tt.input,
					EndTime:    time.Now().Add(remainingDuration),
				}

				utils.TestSetupStatusFile(t, status, file)
				currentStatus := scmd.Status(file.Name(), "notify-send")

				if currentStatus != tt.output {
					t.Errorf(
						"Got %s, want %s",
						currentStatus, tt.output,
					)
				}
			},
		)
	}
}
