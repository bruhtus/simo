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
) {
	var (
		existingEndTime = utils.DetermineEndTime(statusPath)
		isExpired       = time.Now().After(existingEndTime)
	)

	if !isExpired {
		fmt.Printf(
			"Session %s still on going, use reset subcommand to start a new session\n",
			state,
		)
		return
	}

	endTime := time.Now().Add(duration)

	status := utils.Status{
		State:   state,
		EndTime: endTime,
	}

	statusJSON, err := json.Marshal(status)
	utils.CheckError(err)

	utils.WriteStatusFile(statusPath, statusJSON)
}
