package scmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bruhtus/simo/utils"
)

// TODO: add test case.
func OnGoing(
	statusPath string,
	duration time.Duration,
	state utils.StatusState,
	isUseNotify bool,
	isNotifyCmdExist bool,
) {
	var (
		isNotify = false

		currentState    = utils.DetermineStatusState(statusPath)
		existingEndTime = utils.DetermineEndTime(statusPath)
		isExpired       = time.Now().After(existingEndTime)
	)

	if !isExpired {
		fmt.Printf(
			"Session %s still on going, use reset subcommand to start a new session\n",
			currentState,
		)
		return
	}

	if isUseNotify && isNotifyCmdExist {
		isNotify = true
	}

	endTime := time.Now().Add(duration)

	status := utils.Status{
		State:    state,
		IsNotify: isNotify,
		EndTime:  endTime,
	}

	statusJSON, err := json.Marshal(status)
	utils.CheckError(err)

	utils.WriteStatusFile(statusPath, statusJSON)
}
