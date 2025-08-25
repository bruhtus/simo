package scmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bruhtus/simo/utils"
)

// TODO: add test case.
func Pause(statusPath string) {
	var (
		status    = utils.ReadStatusFile(statusPath)
		isExpired = utils.DetermineIsExpired(status.EndTime)

		minutes, seconds = utils.ParseRemainingDuration(
			status.EndTime,
		)
		newPausePoint = fmt.Sprintf("%dm%ds", minutes, seconds)
	)

	if isExpired {
		if status.PausePoint != nil {
			status.PausePoint = nil

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
