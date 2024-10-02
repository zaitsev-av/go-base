package password

import (
	"encoding/json"
	"fmt"
	"go-base/files"
	"go-base/utils"
)

func Password() {
Menu:
	for {
		userOutput := passwordMenu()
		switch userOutput {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			fmt.Println("В разработке")
		default:
			break Menu
		}
	}
}

func findAccount() {
	var outputKey string
	fmt.Println("Введите ключ для поиска")
	fmt.Scanln(&outputKey)

	data := files.ReadFile("accountData.json")

	var jsonData map[string]Account
	err := json.Unmarshal(data, &jsonData)
	result, exist := jsonData[outputKey]

	if !exist {
		fmt.Println("Ключ не найден")
		return
	}

	utils.PrintError(err, "Ошибка парсинга стр 33 findPassword")
	fmt.Println("Login: ", result.Login)
	fmt.Println("Password: ", result.Password)
}
