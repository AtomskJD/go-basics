/**
На вход подается строка. Нужно определить,
является ли она правильной или нет.
Правильная строка начинается с заглавной
буквы и заканчивается точкой.
Если строка правильная - вывести Right
иначе - вывести Wrong
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	a = strings.Trim(a, "\n")
	r := []rune(a)
	t := unicode.IsUpper(r[0])
	t2 := string(r[len(r)-1]) == "."

	if t && t2 {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
