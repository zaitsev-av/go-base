package password

import (
	"errors"
	"fmt"
	"go-base/files"
)

const URL_ERROR = "NO_CORRECT_URL"

func Password() {
	createAccount()
}

func createAccount() {
	login, loginErr := prompt("Введите логин: ")
	if loginErr != nil {
		fmt.Println("Вы не ввели логин")
		return
	}
	password, _ := prompt("Введите пароль: ")
	var passLength string
	if password == "" {
		length, _ := prompt("Введите длинну пароля: ")
		passLength = length
	}
	url, _ := prompt("Введите URL: ")

	userOutput, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Вы ввели некоректный URL")
		return
	}

	userOutput.outputPrompt()
	userOutput.randomPassword(passLength)
	bytsData, err := userOutput.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
	}

	files.WriteFile(bytsData, "accountData.json")

}

func prompt(promptData string) (string, error) {

	var res string
	fmt.Print(promptData)
	fmt.Scanln(&res)

	if res == "" {
		return "", errors.New("NO_OUTPUT")
	}
	return res, nil
}
