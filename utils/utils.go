package utils

import (
	"fmt"
	"go-base/consoleColors"
)

func PrintError(error error, errorText string) {
	colors := consoleColors.Colors()
	if error != nil {
		fmt.Println(colors.Red(errorText), error)
		return
	}
}
