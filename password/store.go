package password

import (
	"encoding/json"
	"fmt"
	"go-base/account"
	"go-base/consoleColors"
	"go-base/utils"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type AccountStore struct {
	Accounts map[string]account.Account `json:"accounts"`
}

type AccountStoreDb struct {
	AccountStore
	db Db
}

func InitializeStore(db Db) *AccountStoreDb {
	data, err := db.Read()
	if err != nil {
		return newStore(db)
	}
	var storeData AccountStoreDb
	err = json.Unmarshal(data, &storeData)
	if err != nil {
		utils.PrintError(err, "Ошибка при декодировании данных -> Unmarshal")
	}
	return &AccountStoreDb{
		AccountStore: AccountStore{
			Accounts: storeData.Accounts,
		},
		db: db,
	}
}

func newStore(db Db) *AccountStoreDb {
	return &AccountStoreDb{
		AccountStore: AccountStore{
			Accounts: make(map[string]account.Account),
		},
		db: db,
	}
}

func (store *AccountStoreDb) AddAccount(key string, data account.Account) {
	// store.Accounts[key] = data
	store.AccountStore.Accounts[key] = data
	dataToBytes, err := json.Marshal(store)
	if err != nil {
		utils.PrintError(err, "Ошибка сириализации данных")
	}
	store.db.Write(dataToBytes)
}

func (store *AccountStoreDb) FindAccount() {
	var outputKey string
	fmt.Println("Введите ключ для поиска")
	fmt.Scanln(&outputKey)

	fmt.Println("Login: ", store.Accounts[outputKey].Login)
	fmt.Println("Password: ", store.Accounts[outputKey].Password)
}

func (store *AccountStoreDb) RemoveAccount() {
	var outputKey string
	fmt.Println("Введите ключ для поиска")
	fmt.Scanln(&outputKey)
	delete(store.Accounts, outputKey)
	file, err := json.Marshal(store)
	if err != nil {
		consoleColors.Colors().Red("Ошибка декодирования")
	}
	store.db.Write(file)
}
