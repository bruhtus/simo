package utils

import (
	"errors"
	"os/exec"
)

func IsNotifyCmdExist(notifyCmd string) bool {
	_, err := exec.LookPath(notifyCmd)

	if errors.Is(err, exec.ErrNotFound) {
		return false
	}

	CheckError(err)
	return true
}

func SendNotify(
	notifyCmd string,
	state StatusState,
) {
	message := "End of pomodoro session!"

	switch state {
	case StateFocus:
		message = "End of focus session!"
	case StateBreak:
		message = "End of break session!"
	}

	args := []string{
		"Pomodoro",
		message,
		"-a",
		"Simo",
		"-u",
		"critical",
	}

	cmd := exec.Command(notifyCmd, args...)

	err := cmd.Run()
	CheckError(err)
}
