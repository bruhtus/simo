package utils

import (
	"time"
)

func ParseDuration(duration string) time.Duration {
	parsedDuration, err := time.ParseDuration(duration)
	CheckError(err)

	return parsedDuration
}
