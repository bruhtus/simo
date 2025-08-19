package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type StatusState string

const (
	StateFocus StatusState = "focus"
	StateBreak StatusState = "break"
)

type Status struct {
	// TODO: add pause, notification.
	State   StatusState `json:"state"`
	EndTime time.Time   `json:"end_time"`
}

func DetermineStateIndicator(state StatusState) string {
	indicator := "undefined"

	switch state {
	case StateFocus:
		indicator = "F"
	case StateBreak:
		indicator = "B"
	}

	return indicator
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

func ReadStatusFile(statusPath string) *Status {
	var (
		status    = new(Status)
		data, err = os.ReadFile(statusPath)
	)

	if !errors.Is(err, os.ErrNotExist) {
		CheckError(err)

		err = json.Unmarshal(data, status)
		CheckError(err)
	}

	return status
}

func DetermineEndTime(statusPath string) time.Time {
	var (
		endTime = time.Now()
		status  = ReadStatusFile(statusPath)
	)

	if status != nil {
		endTime = status.EndTime
	}

	return endTime
}

func DetermineStatusState(statusPath string) StatusState {
	var (
		state  StatusState
		status = ReadStatusFile(statusPath)
	)

	if status != nil {
		state = status.State
	}

	return state
}
