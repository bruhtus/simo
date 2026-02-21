package main

import (
	"flag"
	"fmt"
	"os"

	scmd "github.com/bruhtus/simo/internal"
	"github.com/bruhtus/simo/utils"
)

var (
	focusTime, breakTime               string
	isFocusUseNotify, isBreakUseNotify bool

	statusPath = "/tmp/simo.json"
	notifyCmd  = "notify-send"

	focusCmd = flag.NewFlagSet("focus", flag.ContinueOnError)
	breakCmd = flag.NewFlagSet("break", flag.ContinueOnError)
)

func init() {
	focusCmd.StringVar(
		&focusTime,
		"t",
		"50m",
		"Pomodoro duration time",
	)

	// Use flag `-n` to enable notification.
	focusCmd.BoolVar(
		&isFocusUseNotify,
		"n",
		false,
		"Pomodoro notification",
	)

	breakCmd.StringVar(
		&breakTime,
		"t",
		"10m",
		"Pomodoro duration time",
	)

	// Use flag `-n` to enable notification.
	breakCmd.BoolVar(
		&isBreakUseNotify,
		"n",
		false,
		"Pomodoro notification",
	)
}

func main() {
	if len(os.Args) < 2 {
		utils.HelpUsage()
	}

	isNotifyCmdExist := utils.IsNotifyCmdExist(notifyCmd)

	switch os.Args[1] {
	case "status":
		remaining := scmd.Status(statusPath, notifyCmd)
		fmt.Println(remaining)

	case "focus":
		focusCmd.Parse(os.Args[2:])
		duration := utils.ParseDuration(focusTime)

		scmd.OnGoing(
			statusPath,
			duration,
			utils.StateFocus,
			isFocusUseNotify,
			isNotifyCmdExist,
		)

	case "break":
		breakCmd.Parse(os.Args[2:])
		duration := utils.ParseDuration(breakTime)

		scmd.OnGoing(
			statusPath,
			duration,
			utils.StateBreak,
			isBreakUseNotify,
			isNotifyCmdExist,
		)

	case "reset":
		err := scmd.Reset(statusPath)
		utils.CheckError(err)

	case "pause":
		scmd.Pause(nil, statusPath)

	default:
		utils.HelpUsage()
	}
}
