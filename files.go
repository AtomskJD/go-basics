package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	f, ok := ioutil.ReadFile("readme.md")
	if ok == nil {
		fmt.Printf("%s\n", f)
	}

	// make test data
	dataForFile := []byte("Данные для записи в файл 1234567890!\n")
	if err := ioutil.WriteFile("test.txt", dataForFile, 0755); err != nil {
		fmt.Println("Ошибка записи")
	}

	dataFromFile, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Ошибка чтения")
	} else {
		fmt.Printf("%s\n", dataFromFile)
	}

	// check is data equal
	fmt.Printf("dataForFile == dataFromFile: %v\n",
		bytes.Equal(dataForFile, dataFromFile))

	filesFromDir, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Ошибка чтения директории")
	} else {
		fmt.Println(filesFromDir)
	}

	for _, file := range filesFromDir {
		fmt.Printf("name: %s \t size: %d \t mode: %d\n", file.Name(), file.Size(), file.Mode())
	}
}
