package password

import (
	"fmt"
	"go-base/consoleColors"

	"github.com/fatih/color"
)

func passwordMenu() int {
	colors := consoleColors.Colors()
	var userOutput int
	fmt.Println(colors.YellowBoldUl("1. Создать аккаунт"))
	fmt.Println(colors.YellowBoldUl("2. Найти аккаунт"))
	fmt.Println(colors.YellowBoldUl("3. Посмотреть список ключей"))
	fmt.Println(colors.YellowBoldUl("4. Удалить аккаунт"))
	fmt.Println(colors.YellowBoldUl("5. Выход"))
	fmt.Print(color.HiYellowString("Выберите дальниешее действие: "))
	fmt.Scanln(&userOutput)
	return userOutput
}
