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

	_, err = file.Write(statusJSON)
	if err != nil {
		t.Fatalf(
			"Failed to write data into temporary file: %v",
			err,
		)
	}

	t.Cleanup(func() {
		err := file.Truncate(0)
		if err != nil {
			t.Fatalf(
				"Failed to truncate file: %v",
				err,
			)
		}

		_, err = file.Seek(0, 0) // Seek the beginning of the file.
		if err != nil {
			t.Fatalf(
				"Failed to seek the begining of file: %v",
				err,
			)
		}
	})
}
