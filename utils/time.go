package utils

import (
	"time"
)

func ParseDuration(duration string) time.Duration {
	parsedDuration, err := time.ParseDuration(duration)
	CheckError(err)

	return parsedDuration
}

func ParseRemainingDuration(
	endTime time.Time,
) (time.Duration, time.Duration) {
	remaining := time.
		Until(endTime).
		Round(time.Second)

	minutes := remaining / time.Minute
	remaining -= minutes * time.Minute
	seconds := remaining / time.Second

	return minutes, seconds
}

func DetermineIsExpired(endTime time.Time) bool {
	var (
		isExpired = false

		remainingDuration = time.Until(endTime)
		remainingSeconds  = remainingDuration.
					Round(time.Second).
					Seconds()
	)

	if remainingSeconds <= 0 {
		isExpired = true
	}

	return isExpired
}
