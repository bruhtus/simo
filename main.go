package main

import (
	"flag"
	"fmt"
	"os"

	scmd "github.com/bruhtus/simo/internal"
	"github.com/bruhtus/simo/utils"
)

var (
	t          string
	statusPath = "/tmp/simo.json"

	focusCmd = flag.NewFlagSet("focus", flag.ContinueOnError)
)

func init() {
	focusCmd.StringVar(&t, "t", "50m", "Pomodoro duration time")
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
		duration := utils.ParseDuration(t)
		scmd.Focus(statusPath, duration)

	// TODO: add break subcommand.

	default:
		utils.HelpUsage()
	}
}
