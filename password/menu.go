package password

import (
	"fmt"
	"go-base/consoleColors"
)

func passwordMenu() int {
	colors := consoleColors.Colors()
	var userOutput int
	fmt.Println(colors.YellowBoldUl("1. Создать аккаунт"))
	fmt.Println(colors.YellowBoldUl("2. Найти аккаунт"))
	fmt.Println(colors.YellowBoldUl("3. Удалить аккаунт"))
	fmt.Println(colors.YellowBoldUl("4. Выход"))
	fmt.Scanln(&userOutput)
	return userOutput
}

// func findAccount(accountName string) {

// }
