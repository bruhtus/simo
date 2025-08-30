package scmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bruhtus/simo/utils"
)

func OnGoing(
	statusPath string,
	duration time.Duration,
	state utils.StatusState,
	isUseNotify bool,
	isNotifyCmdExist bool,
) {
	var (
		isNotify  = false
		status    = utils.ReadStatusFile(statusPath)
		isExpired = utils.DetermineIsExpired(status.EndTime)
	)

	if !isExpired {
		fmt.Printf(
			"Session %s still on going, use reset subcommand to start a new session\n",
			status.State,
		)
		return
	}

	if isUseNotify && isNotifyCmdExist {
		isNotify = true
	}

	endTime := time.Now().Add(duration)

	status.State = state
	status.IsNotify = isNotify
	status.EndTime = endTime

	statusJSON, err := json.Marshal(status)
	utils.CheckError(err)

	utils.WriteStatusFile(statusPath, statusJSON)
}
