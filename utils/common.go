package utils

import (
	"fmt"
	"os"
	"runtime"
)

// Credit: inspired by Cobra's CheckErr().
func CheckError(err error) {
	if err != nil {
		customErr := err
		_, file, line, ok := runtime.Caller(1)

		if ok {
			customErr = fmt.Errorf(
				"%s#%d: %v\n",
				file,
				line,
				err,
			)
		}

		fmt.Fprintf(os.Stderr, "Error: %v\n", customErr)
		os.Exit(1)
	}
}

func HelpUsage() {
	fmt.Fprintf(
		os.Stdout,
		"Usage: simo status | focus [-t 10m] \n",
	)
	os.Exit(0)
}
