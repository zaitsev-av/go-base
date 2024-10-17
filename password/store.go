package password

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-base/account"
	"go-base/consoleColors"
	"go-base/encrypter"
	"go-base/utils"
)

var findAccountSubMenu = []string{"По ключу", "По URL", "По логину", "Выберите вариант"}

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type AccountStore struct {
	Accounts map[string]account.Account `json:"accounts"`
}

type AccountStoreDb struct {
	AccountStore
	db        Db
	encrypret encrypter.Encrypter
}

type AccountInfo struct {
	login    string
	password string
}

func updateInfo(login, password string) *AccountInfo {
	return &AccountInfo{
		login:    login,
		password: password,
	}
}

func InitializeStore(db Db, enc encrypter.Encrypter) *AccountStoreDb {
	data, err := db.Read()
	if err != nil {
		return newStore(db, enc)
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
		db:        db,
		encrypret: enc,
	}
}

func newStore(db Db, enc encrypter.Encrypter) *AccountStoreDb {
	return &AccountStoreDb{
		AccountStore: AccountStore{
			Accounts: make(map[string]account.Account),
		},
		db:        db,
		encrypret: enc,
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

func findByKey(store *AccountStoreDb, key string) (*AccountInfo, error) {
	data, ok := store.Accounts[key]

	if !ok {
		return nil, errors.New("NO_ACCOUNTS")
	}
	return updateInfo(data.Login, data.Password), nil

}

// избавится от дублирования (пока не придумал как не изменяя сигнатуру функции)
func findByLogin(store *AccountStoreDb, login string) (*AccountInfo, error) {
	for _, value := range store.Accounts {
		if value.Login == login {
			return updateInfo(value.Login, value.Password), nil
		}
	}
	return nil, errors.New("NO_ACCOUNTS")
}

func findByUrl(store *AccountStoreDb, url string) (*AccountInfo, error) {
	for _, value := range store.Accounts {
		if value.Url == url {
			return updateInfo(value.Login, value.Password), nil
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
