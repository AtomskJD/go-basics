package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	buf := bytes.NewBufferString("11\n22\n33\n")
	in, _ := ioutil.ReadAll(buf)
	var summ int
	for _, el := range in {
		fmt.Println(el)
		sum, _ := strconv.Atoi(string(el))
		summ += sum
	}
	bufOut := bufio.NewWriter(os.Stdout)
	bufOut.WriteString(strconv.Itoa(summ))
	bufOut.Flush()

}
