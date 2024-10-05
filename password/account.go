package password

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-base/files"
	"go-base/utils"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

const URL_ERROR = "NO_CORRECT_URL"

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) randomPassword(passwordLength string) {
	symbols := []rune("qwertyuiop[]asdfghjkl;'zxcvbnm,/1234567890QWERTYUIOPASDFGHJKLZXCVBNM")
	password := ""
	length, _ := strconv.Atoi(passwordLength)
	for i := 0; i < length; i++ {
		password = password + string(symbols[rand.Intn(len(symbols))])
	}
	acc.Password = password
}

func newAccount(login, password, url string) (*Account, error) {

	if !isValidUrl(url) {
		return nil, errors.New(URL_ERROR)
	}

	return &Account{
		Login:     login,
		Password:  password,
		Url:       url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (account *Account) outputPrompt() {
	fmt.Println(*account)
}

func (account *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(account)
	if err != nil {
		return nil, errors.New("ENCODING_ERROR")
	}
	return file, nil
}

func isValidUrl(urls string) bool {
	u, err := url.Parse(urls)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func createAccount() {
	login, loginErr := prompt("Введите логин: ")
	utils.PrintError(loginErr, "Вы не ввели логин")
	password, _ := prompt("Введите пароль: ")
	var passLength string
	if len(password) == 0 {
		length, _ := prompt("Введите длинну пароля: ")
		passLength = length
	}
	url, _ := prompt("Введите URL: ")
	key, _ := prompt("Введите ключ: ")

	userOutput, newAccError := newAccount(login, password, url)
	utils.PrintError(newAccError, "Вы ввели некоректный URL")

	userOutput.outputPrompt()
	if len(userOutput.Password) == 0 {
		userOutput.randomPassword(passLength)
	}
	files.WriteFile(userOutput, "accountData.json", key)

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
