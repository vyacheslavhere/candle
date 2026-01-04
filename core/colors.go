package core

import (
	"os"

	"golang.org/x/term"
)

// Checks colors are enabled
func СolorsEnabled() bool {
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return false
	}

	return term.IsTerminal(int(os.Stdout.Fd()))
}

// Constant colors
const ResetColor = "\033[0m"
const RedColor = "\x1b[38;2;255;64;80m"
const GreenColor = "\x1B[32m"
const YellowColor = "\x1B[38;5;220m"
const BlueColor = "\x1B[34m"
const PurpleColor = "\x1B[35m"
const CyanColor = "\x1B[36m"
const GrayColor = "\x1b[38;5;248m"
const LimeColor = "\x1B[92m"
