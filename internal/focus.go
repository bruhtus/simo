package scmd

import (
	"encoding/json"
	"time"

	"github.com/bruhtus/simo/utils"
)

func Focus(statusPath string, duration time.Duration) {
	endTime := time.Now().Add(duration)

	status := utils.Status{
		State:   utils.StateFocus,
		EndTime: endTime,
	}

	statusJSON, err := json.Marshal(status)
	utils.CheckError(err)

	utils.WriteStatusFile(statusPath, statusJSON)
}
