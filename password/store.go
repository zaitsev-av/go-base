package password

import (
	"encoding/json"
	"fmt"
	"go-base/account"
	"go-base/consoleColors"
	"go-base/utils"
	"os"
)

type AccountStore struct {
	Accounts map[string]account.Account `json:"accounts"`
}

func InitializeStore(data []byte, err error) *AccountStore {
	if err != nil {
		return newStore()
	}
	var storeData AccountStore
	err = json.Unmarshal(data, &storeData)
	if err != nil {
		utils.PrintError(err, "Ошибка при декодировании данных -> Unmarshal")
	}
	return &storeData
}

func newStore() *AccountStore {
	return &AccountStore{
		Accounts: make(map[string]account.Account),
	}
}

func (store *AccountStore) AddAccount(key string, data account.Account) {
	store.Accounts[key] = data
	dataToBytes, err := json.Marshal(store)
	if err != nil {
		utils.PrintError(err, "Ошибка сириализации данных")
	}
	os.WriteFile("accountData.json", dataToBytes, 0644)
}

func (store *AccountStore) FindAccount() {
	var outputKey string
	fmt.Println("Введите ключ для поиска")
	fmt.Scanln(&outputKey)

	fmt.Println("Login: ", store.Accounts[outputKey].Login)
	fmt.Println("Password: ", store.Accounts[outputKey].Password)
}

func (store *AccountStore) RemoveAccount() {
	var outputKey string
	fmt.Println("Введите ключ для поиска")
	fmt.Scanln(&outputKey)
	delete(store.Accounts, outputKey)
	file, err := json.Marshal(store)
	if err != nil {
		consoleColors.Colors().Red("Ошибка декодирования")
	}

	os.WriteFile("accountData.json", file, 0644)
}
