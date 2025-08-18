package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type StatusState string

// Indicator for status command.
const (
	StateFocus StatusState = "F"
	StateBreak StatusState = "B"
)

type Status struct {
	// TODO: add pause, notification.
	State   StatusState `json:"state"`
	EndTime time.Time   `json:"end_time"`
}

func WriteStatusFile(path string, data []byte) {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dir, 0755)
		CheckError(err)
	}

	err := os.WriteFile(path, data, 0644)
	CheckError(err)
}

func DetermineEndTime(statusPath string) time.Time {
	endTime := time.Now()
	data, err := os.ReadFile(statusPath)

	if !errors.Is(err, os.ErrNotExist) {
		CheckError(err)

		status := new(Status)

		err = json.Unmarshal(data, status)
		CheckError(err)

		endTime = status.EndTime
	}

	return endTime
}

func DetermineStatusState(statusPath string) StatusState {
	var state StatusState
	data, err := os.ReadFile(statusPath)

	if !errors.Is(err, os.ErrNotExist) {
		CheckError(err)

		status := new(Status)

		err = json.Unmarshal(data, status)
		CheckError(err)

		state = status.State
	}

	return state
}
