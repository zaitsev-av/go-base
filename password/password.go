package password

import (
	"go-base/account"
	"go-base/encrypter"
	"go-base/files"
	"go-base/utils"

	"github.com/joho/godotenv"
)

const fileName = "accountData.json"

var menuItems = []string{"Создать аккаунт", "Найти аккаунт", "Посмотреть список ключей", "Удалить аккаунт", "Выход", "Выберите дальниешее действие"}

func Password() {
	err := godotenv.Load(".env")
	if err != nil {
		utils.PrintError(err, "Не удалось прочитать .env файл")
	}
	enc := encrypter.NewEncrypter()
	store := InitializeStore(files.NewJsonDb(fileName), *enc)
Menu:
	for {
		userOutput := templateMenu(menuItems)
		switch userOutput {
		case 1:
			key, data := account.CreateAccount()
			store.AddAccount(key, *data)
		case 2:
			store.FindAccount()
		case 3:
			store.KeyList()
		case 4:
			store.RemoveAccount()
		default:
			break Menu
		}
	}
}
