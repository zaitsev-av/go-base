package files

import (
	"encoding/json"
	"fmt"
	"go-base/utils"
	"os"
)

func ReadFile(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(content any, fileName string, key string) {
	openFile, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	utils.PrintError(err, "Ошибка при открытии файла: ")
	defer openFile.Close()

	readFileData, err := os.ReadFile(fileName)
	utils.PrintError(err, "Ошибка при чтении файла: ")
	var jsonData map[string]any
	err = json.Unmarshal(readFileData, &jsonData)
	utils.PrintError(err, "Ошибка парсинга JSON: ")

	jsonData[key] = content

	data, _ := json.Marshal(jsonData)

	if err := openFile.Truncate(0); err != nil {
		fmt.Println("Ошибка при очистке файла:", err)
		return
	}
	if _, err := openFile.Seek(0, 0); err != nil {
		fmt.Println("Ошибка при установке указателя в начало файла:", err)
		return
	}

	if _, err := openFile.Write(data); err != nil {
		fmt.Println("Ошибка записи файла: ", err)
	}
}
