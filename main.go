package main

import (
	"flag"
	"fmt"
	"os"

	scmd "github.com/bruhtus/simo/internal"
	"github.com/bruhtus/simo/utils"
)

var (
	focusTime, breakTime string
	statusPath           = "/tmp/simo.json"

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

	breakCmd.StringVar(
		&breakTime,
		"t",
		"10m",
		"Pomodoro duration time",
	)
}

func main() {
	if len(os.Args) < 2 {
		utils.HelpUsage()
	}

	switch os.Args[1] {
	case "status":
		remaining := scmd.Status(statusPath)
		fmt.Println(remaining)

	case "focus":
		focusCmd.Parse(os.Args[2:])
		duration := utils.ParseDuration(focusTime)
		scmd.OnGoing(statusPath, duration, utils.StateFocus)

	case "break":
		breakCmd.Parse(os.Args[2:])
		duration := utils.ParseDuration(breakTime)
		scmd.OnGoing(statusPath, duration, utils.StateBreak)

	// TODO: add reset subcommand to reset current state.

	default:
		utils.HelpUsage()
	}
}
