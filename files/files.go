package files

import (
	"encoding/json"
	"fmt"
	"go-base/utils"
	"os"
)

func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func WriteFile(content any, fileName string, key string) {
	openFile, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	defer openFile.Close()

	if err != nil {
		if os.IsNotExist(err) {
			file, _ := os.Create(fileName)
			openFile = file
			writeContent := make(map[string]any)
			writeContent[key] = content
			data, _ := json.Marshal(writeContent)
			openFile.Write(data)
			return
		}
	}

	utils.PrintError(err, "Ошибка при открытии файла: ")

	readFileData, err := os.ReadFile(fileName)
	utils.PrintError(err, "Ошибка при чтении файла: ")
	var jsonData map[string]any
	err = json.Unmarshal(readFileData, &jsonData)
	utils.PrintError(err, "Ошибка парсинга JSON: ")

	jsonData[key] = content

	data, _ := json.Marshal(jsonData)

	if _, err := openFile.Write(data); err != nil {
		fmt.Println("Ошибка записи файла: ", err)
	}
}

func RemoveAccount() {
	var outputKey string
	fmt.Println("Введите ключ для поиска")
	fmt.Scanln(&outputKey)

	openFile, err := os.OpenFile("accountData.json", os.O_RDWR, 0644)
	defer openFile.Close()
	utils.PrintError(err, "Ошибка при открытии файла: ")
	readFileData, err := os.ReadFile("accountData.json")
	utils.PrintError(err, "Ошибка при чтении файла: ")
	var jsonData map[string]any
	err = json.Unmarshal(readFileData, &jsonData)
	utils.PrintError(err, "Ошибка парсинга JSON: ")

	delete(jsonData, outputKey)
}
