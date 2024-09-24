package password

import (
	"fmt"
	"math/rand"
)

type account struct {
	login    string
	password string
	url      string
}

func (account *account) outputPrompt() {
	fmt.Println(*account)
}

func Password() {
	length := promptNum("Введите длинну пароля: ")
	randomPassword(length)
	//-----------------------
	login := prompt("Введите логин: ")
	password := prompt("Введите пароль: ")
	url := prompt("Введите URL: ")

	var userOutput = account{
		login,
		password,
		url,
	}

	userOutput.outputPrompt()
	fmt.Println(userOutput, "userOutput")
}

func prompt(promptData string) string {

	var res string

	fmt.Print(promptData)
	fmt.Scan(&res)

	return res
}

func promptNum(promptData string) int {

	var res int

	fmt.Print(promptData)
	fmt.Scan(&res)

	return res
}

func randomPassword(passwordLength int) string {
	symbols := []rune("qwertyuiop[]asdfghjkl;'zxcvbnm,/1234567890QWERTYUIOPASDFGHJKLZXCVBNM")
	password := ""
	for i := 0; i < passwordLength; i++ {
		password = password + string(symbols[rand.Intn(len(symbols))])

	}
	fmt.Println(password)
	return ""
}
