package scmd

import (
	"encoding/json"
	"time"

	"github.com/bruhtus/simo/utils"
)

func OnGoing(
	statusPath string,
	duration time.Duration,
	state utils.StatusState,
) {
	// TODO:
	// abort when we're already in on going state,
	// not yet expired state.
	endTime := time.Now().Add(duration)

	status := utils.Status{
		State:   state,
		EndTime: endTime,
	}

	statusJSON, err := json.Marshal(status)
	utils.CheckError(err)

	utils.WriteStatusFile(statusPath, statusJSON)
}
