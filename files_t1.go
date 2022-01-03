package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func main() {
	buf := bytes.NewBufferString("11\n22\n33\n")
	in := bufio.NewScanner(buf)
	var summ int
	for in.Scan() {
		sum, _ := strconv.Atoi(in.Text())
		summ += sum
	}
	bufOut := bufio.NewWriter(os.Stdout)
	bufOut.WriteString(strconv.Itoa(summ))
	bufOut.Flush()

}
