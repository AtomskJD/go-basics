package main

import "os"

func main() {
	file1, _ := os.Create("os_test.txt")
	file1.WriteString("Строка 1 \n")
	file1.WriteString("Строка 2 \n")
	os.Rename("os_test.txt", "os_test_renamed.txt")
	file1.Close()
}
