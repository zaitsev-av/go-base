package finance

import "fmt"

func Finance() {
	transactions := []int{5, 10, -7}
	fmt.Println("Давайте соберем все ваши транзакции")
	for {
		var transaction int
		fmt.Print("Введите сумму транзакции: ")
		fmt.Scan(&transaction)
		transactions = append(transactions, transaction)
		var response string
		fmt.Println("Есть еще транзакции? (y/n)")
		fmt.Scan(&response)
		if response == "y" || response == "Y" || response == "yes" {
			continue
		}
		break
	}

	result := calculateBalance(transactions)

	fmt.Println("Ваш баланс составляет: ", result)

}

func calculateBalance(transactions []int) (result int) {
	for _, transaction := range transactions {
		result += transaction
	}
	return result
}
