package password

import (
	"go-base/account"
	"go-base/files"
)

const fileName = "accountData.json"

func Password() {
	db := files.NewJsonDb(fileName)
	data, err := db.Read()
	store := InitializeStore(data, err)

Menu:
	for {
		userOutput := passwordMenu()
		switch userOutput {
		case 1:
			key, data := account.CreateAccount()
			store.AddAccount(key, *data)
		case 2:
			store.FindAccount()
		case 3:
			store.RemoveAccount()
		default:
			break Menu
		}
	}
}
