package scmd_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	scmd "github.com/bruhtus/simo/internal"
	"github.com/bruhtus/simo/utils"
)

func TestOnGoingEndTime(t *testing.T) {
	durationCases := []time.Duration{
		time.Duration(0 * time.Second),
		time.Duration(1 * time.Second),
		time.Duration(1 * time.Minute),
		time.Duration(1 * time.Hour),
	}

	endTimeCases := []struct {
		duration time.Duration
		output   time.Time
	}{
		{
			durationCases[0],
			time.
				Now().
				Add(durationCases[0]).
				Round(time.Second),
		},
		{
			durationCases[1],
			time.
				Now().
				Add(durationCases[1]).
				Round(time.Second),
		},
		{
			durationCases[2],
			time.
				Now().
				Add(durationCases[2]).
				Round(time.Second),
		},
		{
			durationCases[3],
			time.
				Now().
				Add(durationCases[3]).
				Round(time.Second),
		},
	}

	t.Run("Status file exist", func(t *testing.T) {
		var (
			dirPath = t.TempDir()
			file    = utils.TestSetupTempFile(t, dirPath)
		)

		t.Cleanup(func() {
			err := file.Close()
			if err != nil {
				t.Fatalf(
					"Failed to close file: %v",
					err,
				)
			}
		})

		for _, tt := range endTimeCases {
			t.Run(
				tt.duration.String(),
				func(t *testing.T) {
					status := utils.Status{
						State:      utils.StateFocus,
						IsNotify:   true,
						PausePoint: nil,
					}

					utils.TestSetupStatusFile(t, status, file)

					scmd.OnGoing(
						file.Name(),
						tt.duration,
						utils.StateFocus,
						true,
						true,
					)

					resultJSON := utils.ReadStatusFile(file.Name())
					endTime := resultJSON.EndTime.Round(time.Second)

					comparation := endTime.Compare(tt.output)

					if comparation != 0 {
						t.Errorf(
							"Got %v, want %v. Compare result: %d",
							endTime,
							tt.output,
							comparation,
						)
					}
				},
			)
		}
	})

	t.Run("Status file not exist", func(t *testing.T) {
		dirPath := t.TempDir()

		filePath := fmt.Sprintf(
			"%s/simo-%d.json",
			dirPath,
			time.Now().UnixMilli(),
		)

		for _, tt := range endTimeCases {
			t.Run(
				tt.duration.String(),
				func(t *testing.T) {
					_, err := os.Stat(filePath)

					if err != nil && !errors.Is(err, os.ErrNotExist) {
						t.Fatalf(
							"Error read status file: %v",
							err,
						)
					}

					scmd.OnGoing(
						filePath,
						tt.duration,
						utils.StateFocus,
						true,
						true,
					)

					resultJSON := utils.ReadStatusFile(filePath)
					endTime := resultJSON.EndTime.Round(time.Second)

					t.Cleanup(func() {
						// Reset the end time so that it become expired.
						if !errors.Is(err, os.ErrNotExist) {
							resultJSON.EndTime = time.Now()
							statusJSON, err := json.Marshal(resultJSON)

							if err != nil {
								t.Fatalf(
									"Error json marshal existing status file: %v",
									err,
								)
							}

							utils.WriteStatusFile(filePath, statusJSON)
						}
					})

					comparation := endTime.Compare(tt.output)

					if comparation != 0 {
						t.Errorf(
							"Got %v, want %v. Compare result: %d",
							endTime,
							tt.output,
							comparation,
						)
					}
				},
			)
		}
	})
}

func TestOnGoingIsNotify(t *testing.T) {
	isNotifyCases := []struct {
		inputUseNotify      bool
		inputNotifyCmdExist bool
		output              bool
	}{
		{true, true, true},
		{false, true, false},
		{true, false, false},
		{false, false, false},
	}

	var (
		dirPath = t.TempDir()
		file    = utils.TestSetupTempFile(t, dirPath)
	)

	t.Cleanup(func() {
		err := file.Close()
		if err != nil {
			t.Fatalf(
				"Failed to close file: %v",
				err,
			)
		}
	})

	for _, tt := range isNotifyCases {
		t.Run(
			fmt.Sprintf(
				"isUseNotify %t, isNotifyCmdExist %t",
				tt.inputUseNotify, tt.inputNotifyCmdExist,
			),
			func(t *testing.T) {
				status := utils.Status{
					State:      utils.StateFocus,
					IsNotify:   true,
					PausePoint: nil,
				}

				utils.TestSetupStatusFile(t, status, file)

				scmd.OnGoing(
					file.Name(),
					time.Duration(0*time.Second),
					utils.StateFocus,
					tt.inputUseNotify,
					tt.inputNotifyCmdExist,
				)

				resultJSON := utils.ReadStatusFile(file.Name())

				if resultJSON.IsNotify != tt.output {
					t.Errorf(
						"Got %v, want %v",
						resultJSON.IsNotify,
						tt.output,
					)
				}
			},
		)
	}
}
