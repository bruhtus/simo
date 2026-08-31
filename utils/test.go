package utils

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSetupTempFile(
	t *testing.T,
	dirPath string,
) *os.File {
	t.Helper()

	file, err := os.CreateTemp(dirPath, "simo-*.json")
	if err != nil {
		t.Fatalf(
			"Failed to create temporary file: %v",
			err,
		)
	}

	return file
}

func TestSetupStatusFile(
	t *testing.T,
	status Status,
	file *os.File,
) {
	t.Helper()

	statusJSON, err := json.Marshal(status)
	if err != nil {
		t.Fatalf(
			"Failed to marshal status json: %v",
			err,
		)
	}

	// Looks like there's a difference between using (*os.File).Write() and
	// (*os.File).Truncate(0) with os.WriteFile().
	// If we use (*os.File).Write() and (*os.File).Truncate(), there's a
	// chance that the file content become stale (?), which can cause problem
	// when we require the test file to use new data.
	err = os.WriteFile(file.Name(), statusJSON, 0644)
	if err != nil {
		t.Fatalf(
			"Failed to write data into temporary file: %v",
			err,
		)
	}
}
