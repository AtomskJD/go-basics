// Password checker
package main

import (
	"fmt"
	"unicode"
)

// check if string more than five characters
func isFive(s string) bool {
	r := []rune(s)
	if len(r) >= 5 {
		return true
	}
	return false
}

// check if is only latin characters in string
func isLatin(s string) bool {
	r := []rune(s)
	for _, el := range r {
		if !(unicode.Is(unicode.Latin, el) || unicode.IsDigit(el)) {
			return false
		}
	}
	return true
}
func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("Проверьте типы входных параметров")
		return
	}
	if isFive(s) && isLatin(s) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
