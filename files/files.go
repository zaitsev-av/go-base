package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(content, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println("Success")
	file.Close()
}
