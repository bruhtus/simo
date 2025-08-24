package scmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bruhtus/simo/utils"
)

// TODO: add test case.
// FIXME: should not be paused when remaining duration is 0 seconds.
func Pause(statusPath string) {
	var (
		status = utils.ReadStatusFile(statusPath)

		existingEndTime = utils.DetermineEndTime(statusPath)
		isExpired       = time.Now().After(existingEndTime)

		minutes, seconds = utils.ParseRemainingDuration(
			existingEndTime,
		)
		newPausePoint = fmt.Sprintf("%dm%ds", minutes, seconds)
	)

	if isExpired {
		if status.PausePoint != nil {
			status.EndTime = existingEndTime

			statusJSON, err := json.Marshal(status)
			utils.CheckError(err)

			utils.WriteStatusFile(statusPath, statusJSON)
		}
		return
	}

	if status.PausePoint != nil {
		remainingDuration := utils.ParseDuration(*status.PausePoint)
		newEndTime := time.Now().Add(remainingDuration)

		status.PausePoint = nil
		status.EndTime = newEndTime

		statusJSON, err := json.Marshal(status)
		utils.CheckError(err)

		utils.WriteStatusFile(statusPath, statusJSON)
		return
	}

	status.PausePoint = &newPausePoint

	statusJSON, err := json.Marshal(status)
	utils.CheckError(err)

	utils.WriteStatusFile(statusPath, statusJSON)
}
