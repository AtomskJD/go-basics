package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("files/task.data")
	rd := bufio.NewReader(file)
	i := 1
	for s, err := rd.ReadString(';'); err == nil && err != io.EOF; {
		if s == "114969837749034510;" {
			fmt.Printf("%s\n", s)
			fmt.Println(i)
		}
		s, err = rd.ReadString(';')
		i++
	}
}
