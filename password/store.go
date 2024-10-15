package password

import (
	"encoding/json"
	"errors"
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

type AccountInfo struct {
	login    string
	password string
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
	store.Accounts[key] = data
	dataToBytes, err := json.Marshal(store)
	if err != nil {
		utils.PrintError(err, "Ошибка сириализации данных")
	}
	store.db.Write(dataToBytes)
}

var findAccountSubMenu = []string{"По ключу", "По URL", "По логину", "Выберите вариант"}

func findByKey1(store *AccountStoreDb) {
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

func findByKey(store *AccountStoreDb, key string) (*AccountInfo, error) {
	data, ok := store.Accounts[key]

	if !ok {
		return nil, errors.New("NO_ACCOUNTS")
	}
	return &AccountInfo{
		login:    data.Login,
		password: data.Password,
	}, nil

}

func findByLogin(store *AccountStoreDb, login string) (*AccountInfo, error) {
	for _, value := range store.Accounts {
		if value.Login == login {
			return &AccountInfo{
				login:    value.Login,
				password: value.Password,
			}, nil
		}
	}
	return nil, errors.New("NO_ACCOUNTS")
}

func findByUrl(store *AccountStoreDb, url string) (*AccountInfo, error) {
	for _, value := range store.Accounts {
		if value.Url == url {
			return &AccountInfo{
				login:    value.Login,
				password: value.Password,
			}, nil
		}
	}
	return nil, errors.New("NO_ACCOUNTS")
}

var findAccountFuncs = map[int]func(store *AccountStoreDb, findParams string) (*AccountInfo, error){
	1: findByKey,
	2: findByUrl,
	3: findByLogin,
}

func (store *AccountStoreDb) FindAccount() {
	userOutput := templateMenu(findAccountSubMenu)
	var outputKey string
	fmt.Print("Введите значени для поиска: ")
	fmt.Scanln(&outputKey)
	findFunc := findAccountFuncs[userOutput]
	data, err := findFunc(store, outputKey)
	// data, ok := store.Accounts[outputKey]

	if err != nil {
		fmt.Println(consoleColors.Colors().Red("По данному ключу ничего не найдено, проверьте ключ"))
		return
	}

	fmt.Println(consoleColors.Colors().Success("Login: ", data.login))
	fmt.Println(consoleColors.Colors().Success("Password: ", data.password))
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
