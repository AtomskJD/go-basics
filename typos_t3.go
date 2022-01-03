package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, err := csv.NewReader(os.Stdin).Read()
	fmt.Println(err)
	if err == nil {
		fmt.Println(a, err)
		a1 := strings.Join(a, ".")
		a1 = strings.Replace(a1, " ", "", -1)
		arr := strings.Split(a1, ";")
		fl1, _ := strconv.ParseFloat(arr[0], 64)
		fl2, _ := strconv.ParseFloat(arr[1], 64)
		if fl2 != 0 {
			res := fl1 / fl2
			s := strconv.FormatFloat(res, 'f', 4, 64)
			fmt.Println(s)
		}
	}
}
