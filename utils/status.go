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
	State    StatusState `json:"state"`
	IsNotify bool        `json:"is_notify"`

	// Duration left where we hit pause.
	PausePoint *string   `json:"pause_point"`
	EndTime    time.Time `json:"end_time"`
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

// To prevent race condition when another process (e.g simo-gui) trying to
// access simo file. Due to wrong timing, when we write the data to simo file
// directly, another process might be in the middle of reading the file and
// might throw error "unexpected end of JSON input" when unmarshal the
// incomplete json file.
// Reference:
// https://dev.to/catatsuy/safely-updating-existing-files-in-go-1hlc
func WriteStatusFile(path string, data []byte) {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dir, 0755)
		CheckError(err)
	}

	tmp, err := os.CreateTemp(dir, "simo-*.json")
	CheckError(err)

	// Clean up if there's an error.
	defer os.Remove(tmp.Name())

	_, err = tmp.Write(data)
	CheckError(err)

	err = tmp.Close()
	CheckError(err)

	err = os.Rename(tmp.Name(), path)
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
