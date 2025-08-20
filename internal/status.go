package scmd

import (
	"fmt"
	"time"

	"github.com/bruhtus/simo/utils"
)

// TODO: add test case.
func Status(statusPath string) string {
	var (
		defaultRemaining = "--:--"
		endTime          = utils.DetermineEndTime(statusPath)
		isExpired        = time.Now().After(endTime)
	)

	var (
		state      = utils.DetermineStatusState(statusPath)
		indicator  = utils.DetermineStateIndicator(state)
		pausePoint = utils.DeterminePausePoint(statusPath)
	)

	if pausePoint != nil {
		remainingDuration := utils.ParseDuration(*pausePoint)

		minutes := remainingDuration / time.Minute
		remainingDuration -= minutes * time.Minute
		seconds := remainingDuration / time.Second

		remainingString := fmt.Sprintf(
			"P%s%02d:%02d",
			indicator,
			minutes,
			seconds,
		)

		return remainingString
	}

	if isExpired {
		return defaultRemaining
	}

	minutes, seconds := utils.ParseRemainingDuration(endTime)

	remainingString := fmt.Sprintf(
		"%s%02d:%02d",
		indicator,
		minutes,
		seconds,
	)

	return remainingString
}
