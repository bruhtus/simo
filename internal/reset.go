package scmd

import (
	"errors"
	"os"
)

// TODO: add test case.
func Reset(statusPath string) error {
	err := os.Remove(statusPath)

	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return err
}
