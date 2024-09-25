package password

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) randomPassword(passwordLength string) {
	symbols := []rune("qwertyuiop[]asdfghjkl;'zxcvbnm,/1234567890QWERTYUIOPASDFGHJKLZXCVBNM")
	password := ""
	length, _ := strconv.Atoi(passwordLength)
	for i := 0; i < length; i++ {
		password = password + string(symbols[rand.Intn(len(symbols))])
	}
	acc.password = password
}

func newAccount(login, password, url string) (*account, error) {

	if !isValidUrl(url) {
		return nil, errors.New(URL_ERROR)
	}

	return &account{
		login:    login,
		password: password,
		url:      url,
	}, nil
}

func (account *account) outputPrompt() {
	fmt.Println(*account)
}

func isValidUrl(urls string) bool {
	u, err := url.Parse(urls)
	return err == nil && u.Scheme != "" && u.Host != ""
}
