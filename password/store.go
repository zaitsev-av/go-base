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
	fmt.Print("Введите ключ для поиска: ")
	fmt.Scanln(&outputKey)

	data, ok := store.Accounts[outputKey]

	if !ok {
		fmt.Println(consoleColors.Colors().Red("По данному ключу ничего не найдено, проверьте ключ"))
		return
	}
	fmt.Println(consoleColors.Colors().Success("Login: ", data.Login))
	fmt.Println(consoleColors.Colors().Success("Password: ", data.Password))
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

func (store *AccountStoreDb) KeyList() {
	for key := range store.Accounts {
		fmt.Printf("|%6s	|	%6s|\n", key, store.Accounts[key].Url)
	}
}
