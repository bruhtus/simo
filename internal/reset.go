package scmd

import (
	"errors"
	"os"
)

func Reset(statusPath string) error {
	err := os.Remove(statusPath)

	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return err
}
