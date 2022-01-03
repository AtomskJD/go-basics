package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	ExampleFunctionName()

	// анонимная функция с присвоением переменной
	fn := func(a, b int) int { return a + b }
	fmt.Println(fn(2, 2))

	// анонимная функция с последующим запуском
	func(a, b int) {
		fmt.Println(a, b)
	}(22, 33)

	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Println(x)
	fmt.Printf("%T", x)

}

func ExampleFunctionName() {
	src := "aBcDeFg"
	test := "AbCdEfG"

	src = strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		} else {
			return unicode.ToLower(r)
		}
	}, src)
	fmt.Printf("Инвертированная строка: %s. Результат: %v.\n", src, src == test)
}
