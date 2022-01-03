package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func mywalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if strings.Contains(info.Name(), "bin") {
		// fmt.Printf("%d ", len(info.Name())+1)
		fmt.Printf("%v ", info.Name())
		return nil
	} else {
		return nil
	}
}

func main() {
	const root = "."
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
}
