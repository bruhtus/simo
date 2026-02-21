package scmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bruhtus/simo/utils"
)

func Pause(
	now func() time.Time,
	statusPath string,
) {
	// Substitute this in test.
	// Reference:
	// https://stackoverflow.com/a/25791617
	if now == nil {
		now = time.Now
	}

	var (
		status    = utils.ReadStatusFile(statusPath)
		isExpired = utils.DetermineIsExpired(status.EndTime)

		minutes, seconds = utils.ParseRemainingDuration(
			status.EndTime,
		)
		newPausePoint = fmt.Sprintf("%dm%ds", minutes, seconds)
	)

	if status.PausePoint != nil {
		remainingDuration := utils.ParseDuration(*status.PausePoint)
		newEndTime := now().Add(remainingDuration)

		status.PausePoint = nil
		status.EndTime = newEndTime

		statusJSON, err := json.Marshal(status)
		utils.CheckError(err)

		utils.WriteStatusFile(statusPath, statusJSON)
		return
	}

	if isExpired {
		return
	}

	status.PausePoint = &newPausePoint

	statusJSON, err := json.Marshal(status)
	utils.CheckError(err)

	utils.WriteStatusFile(statusPath, statusJSON)
}
