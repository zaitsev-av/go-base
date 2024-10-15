package password

import (
	"fmt"
	"go-base/consoleColors"

	"github.com/fatih/color"
)

func templateMenu(items []string) int {
	colors := consoleColors.Colors()
	var userOutput int
	for i, value := range items {
		if len(items)-1 != i {
			message := fmt.Sprintf("%d. %s", i+1, value)
			fmt.Println(colors.YellowBoldUl(message))
		} else {
			message := fmt.Sprintf("%s: ", value)
			fmt.Print(color.HiYellowString(message))
		}
	}
	fmt.Scanln(&userOutput)
	return userOutput
}
