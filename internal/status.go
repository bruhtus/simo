package scmd

import (
	"encoding/json"
	"fmt"

	"github.com/bruhtus/simo/utils"
)

// TODO: add test case.
func Status(
	statusPath string,
	notifyCmd string,
) string {
	var (
		defaultRemaining = "--:--"
		status           = utils.ReadStatusFile(statusPath)
		isExpired        = utils.DetermineIsExpired(status.EndTime)
		indicator        = utils.DetermineStateIndicator(status.State)
	)

	if status.PausePoint != nil {
		remainingDuration := utils.ParseDuration(
			*status.PausePoint,
		)

		minutes, seconds := utils.GetDurationMinutesAndSeconds(
			remainingDuration,
		)

		remainingString := fmt.Sprintf(
			"P%s%02d:%02d",
			indicator,
			minutes,
			seconds,
		)

		return remainingString
	}

	if isExpired {
		statusFile := utils.ReadStatusFile(statusPath)

		if statusFile.IsNotify {
			utils.SendNotify(notifyCmd, status.State)
			statusFile.IsNotify = false

			statusJSON, err := json.Marshal(statusFile)
			utils.CheckError(err)

			utils.WriteStatusFile(statusPath, statusJSON)
		}

		return defaultRemaining
	}

	minutes, seconds := utils.ParseRemainingDuration(
		status.EndTime,
	)

	remainingString := fmt.Sprintf(
		"%s%02d:%02d",
		indicator,
		minutes,
		seconds,
	)

	return remainingString
}
