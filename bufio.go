package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Ошибка")
	}
	defer file.Close()
	rd := bufio.NewReader(file)

	buf := make([]byte, 10)
	n, err := rd.Read(buf)

	if err != nil && err != io.EOF {
		fmt.Println("Ошибка или конец файла")
	}
	fmt.Printf("Считан буфер на %d байт: %s\n", n, buf)

	s, err := rd.ReadString('\n')
	fmt.Println(s)

	file1, err := os.Open("readme.md")
	if err != nil {
		fmt.Println("Ошибка")
	}
	defer file.Close()

	sc := bufio.NewScanner(file1)

	// Я заранее записал в файл 5 цифр, каждую на новой строке
	for sc.Scan() { // возвращает true, пока файл не будет прочитан до конца
		fmt.Printf("%s\n", sc.Text())

	}
}
