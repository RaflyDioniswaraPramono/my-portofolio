package utils

import "github.com/fatih/color"

func SuccessCLIMessage(message string) {
	color.Green(message)
}

func WarningCLIMessage(message string) {
	color.Yellow(message)
}

func FailedCLIMessage(message string) {
	color.Red(message)
}
