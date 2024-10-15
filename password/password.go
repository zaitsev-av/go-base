package password

import (
	"go-base/account"
	"go-base/files"
)

const fileName = "accountData.json"

var menuItems = []string{"Создать аккаунт", "Найти аккаунт", "Посмотреть список ключей", "Удалить аккаунт", "Выход", "Выберите дальниешее действие"}

func Password() {
	store := InitializeStore(files.NewJsonDb(fileName))
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
