package scmd

import (
	"fmt"
	"time"

	"github.com/bruhtus/simo/utils"
)

func Status(statusPath string) string {
	var (
		defaultRemaining = "--:--"
		endTime          = utils.DetermineEndTime(statusPath)
		isExpired        = time.Now().After(endTime)
	)

	var (
		state     = utils.DetermineStatusState(statusPath)
		indicator = utils.DetermineStateIndicator(state)
	)

	if isExpired {
		return defaultRemaining
	}

	remaining := time.
		Until(endTime).
		Round(time.Second)

	minutes := remaining / time.Minute
	remaining -= minutes * time.Minute
	seconds := remaining / time.Second

	remainingString := fmt.Sprintf(
		"%s%02d:%02d",
		indicator,
		minutes,
		seconds,
	)

	return remainingString
}
