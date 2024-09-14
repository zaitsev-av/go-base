package finance

import (
	"fmt"
	"go-base/consoleColors"
)

func Finance() {
	colors := consoleColors.Colors()
	transactions := []int{5, 10, -7}
	fmt.Print(colors.YellowBoldUl(colors.FgCyan("                                                                               Your tranzations     ( ͡👁️ ͜ʖ ͡👁️)                                                                         ")))
	fmt.Println("")
	fmt.Print(colors.RedBold(` 
										█▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█	
										█░░╦─╦╔╗╦─╔╗╔╗╔╦╗╔╗░░█	
										█░░║║║╠─║─║─║║║║║╠─░░█	
										█░░╚╩╝╚╝╚╝╚╝╚╝╩─╩╚╝░░█	
										█▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█
`))

	fmt.Println(colors.YellowBoldUl("Давайте соберем все ваши транзакции"))
	for {
		var transaction int
		fmt.Print(colors.FgCyan("Введите сумму транзакции: "))
		fmt.Scan(&transaction)
		transactions = append(transactions, transaction)
		var response string
		fmt.Println(colors.Yellow("Есть еще транзакции? (y/n)"))
		fmt.Scan(&response)
		if response == "y" || response == "Y" || response == "yes" {
			continue
		}
		break
	}

	result := calculateBalance(transactions)

	fmt.Println(colors.Success("Ваш баланс составляет: ", result))
}
func calculateBalance(transactions []int) (result int) {
	for _, transaction := range transactions {
		result += transaction
	}
	return result
}
