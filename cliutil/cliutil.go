package cliutil

import (
	"fmt"
	"strings"
)

func Prompt(format string, a ...interface{}) bool {
	if !strings.Contains(format, "[y/n]") {
		format += " [y/n] "
	}
	if len(a) == 0 {
		fmt.Print(format)
	} else {
		fmt.Printf(format, a...)
	}
	return handlePromptResponse()
}

func handlePromptResponse() bool {
	var response string
	_, err := fmt.Scanln(&response)

	if err != nil {
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))

	if response == "y" || response == "yes" {
		return true
	}

	return false
}
