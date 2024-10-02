package password

import "fmt"

func Password() {
Menu:
	for {
		userOutput := passwordMenu()
		switch userOutput {
		case 1:
			createAccount()
		case 2:
			fmt.Println("В разработке")
		case 3:
			fmt.Println("В разработке")
		default:
			break Menu
		}
	}
}
