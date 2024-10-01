package password

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

type account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *account) randomPassword(passwordLength string) {
	symbols := []rune("qwertyuiop[]asdfghjkl;'zxcvbnm,/1234567890QWERTYUIOPASDFGHJKLZXCVBNM")
	password := ""
	length, _ := strconv.Atoi(passwordLength)
	for i := 0; i < length; i++ {
		password = password + string(symbols[rand.Intn(len(symbols))])
	}
	acc.Password = password
}

func newAccount(login, password, url string) (*account, error) {

	if !isValidUrl(url) {
		return nil, errors.New(URL_ERROR)
	}

	return &account{
		Login:     login,
		Password:  password,
		Url:       url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (account *account) outputPrompt() {
	fmt.Println(*account)
}

func (account *account) ToBytes() ([]byte, error) {
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
