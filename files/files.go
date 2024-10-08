package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		fileName: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

// func WriteFile[T any](content T, fileName string, key string) {
// 	openFile, err := os.OpenFile(fileName, os.O_RDWR, 0644) //os.O_RDWR|os.O_CREATE можно использовать флаг os.O_CREATE, но нужно подругому обработать запись в данный файл
// 	defer openFile.Close()

// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			file, _ := os.Create(fileName)
// 			openFile = file
// 			writeContent := make(map[string]T)
// 			writeContent[key] = content
// 			data, _ := json.Marshal(writeContent)
// 			openFile.Write(data)
// 			return
// 		}
// 	}

// 	utils.PrintError(err, "Ошибка при открытии файла: ")

// 	readFileData, err := os.ReadFile(fileName)
// 	utils.PrintError(err, "Ошибка при чтении файла: ")
// 	var jsonData map[string]any
// 	err = json.Unmarshal(readFileData, &jsonData)
// 	utils.PrintError(err, "Ошибка парсинга JSON: ")

// 	jsonData[key] = content

// 	data, _ := json.Marshal(jsonData)

// 	if _, err := openFile.Write(data); err != nil {
// 		fmt.Println("Ошибка записи файла: ", err)
// 	}
// }
