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
			time.Duration(-1 * time.Second), "",
			false, false,
		},
		{
			time.Duration(0 * time.Second), "",
			false, false,
		},
		{
			time.Duration(1 * time.Second), "0m1s",
			true, true,
		},
		{
			time.Duration(1 * time.Minute), "1m0s",
			false, false,
		},
		{
			time.Duration(60 * time.Minute), "60m0s",
			true, true,
		},
		{
			time.Duration(90 * time.Minute), "90m0s",
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
				scmd.Pause(nil, file.Name())

				resultJSON := utils.ReadStatusFile(file.Name())

				if resultJSON.PausePoint == nil &&
					tt.remainingDurationOutput != "" {
					t.Errorf("Got unexpected nil")
				}

				if resultJSON.PausePoint != nil &&
					*resultJSON.PausePoint != tt.remainingDurationOutput {
					t.Errorf(
						"Got %s, want %s",
						*resultJSON.PausePoint,
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

func TestResume(t *testing.T) {
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
		pausePoint      string
		endTime         time.Duration
		expectedEndTime time.Duration
	}{
		{
			"0m1s",
			time.Duration(-1 * time.Second),
			time.Duration(1 * time.Second),
		},
		{
			"0m10s",
			time.Duration(-10 * time.Second),
			time.Duration(10 * time.Second),
		},
		{
			"1m0s",
			time.Duration(-1 * time.Minute),
			time.Duration(1 * time.Minute),
		},
		{
			"60m0s",
			time.Duration(-60 * time.Minute),
			time.Duration(60 * time.Minute),
		},
		{
			"90m0s",
			time.Duration(-90 * time.Minute),
			time.Duration(90 * time.Minute),
		},
	}

	for _, tt := range remainingDurationCases {
		t.Run(
			tt.endTime.String(),
			func(t *testing.T) {
				endTime := time.Now().Add(tt.endTime)
				expectedEndTime := endTime.Add(tt.expectedEndTime)

				status := utils.Status{
					State:      utils.StateFocus,
					IsNotify:   false,
					PausePoint: &tt.pausePoint,
					EndTime:    endTime,
				}

				utils.TestSetupStatusFile(t, status, file)

				// Substitute the time.Now() result.
				now := func() time.Time {
					return endTime
				}

				scmd.Pause(now, file.Name())
				resultJSON := utils.ReadStatusFile(file.Name())

				comparation := expectedEndTime.Compare(resultJSON.EndTime)

				if comparation != 0 {
					t.Errorf(
						"Got %v, want %v",
						endTime,
						expectedEndTime,
					)
				}
			},
		)
	}
}
