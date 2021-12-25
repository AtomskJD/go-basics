package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Создадим строковый литерал s, значение которого "Это строка".
	// Строка состоит из 10 символов.
	var s string = "Это строка"

	// Однако длина строки len(s) составит 19 байт, т.к. использованные кирилические символы
	// занимают 2 байта, а пробел занимает 1 байт.
	fmt.Printf("Длина строки: %d байт\n", len(s))

	// Посмотрим как взять подстроку
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\"\n", s[7:])

	/*
		Попробуем изменить что-то встроке:
		s[3] = 12
		Ошибка компиляции: cannot assign to s[3], потому что строки - неизменяемые последовательности.
	*/

	// "Изменим строку", создав новую
	s = s + " Новая строка"
	fmt.Printf("%v\n", s)

	// А теперь проитерируемся по этой строке
	for _, b := range s {
		fmt.Printf("%c ", b)
	}
	fmt.Print("\n")

	test := "Тестер тестировлвл"
	fmt.Println(strings.Count(test, "е"))
	fmt.Println(strings.Contains(test, "ст"))
	fmt.Println(strings.HasPrefix(test, "Тес"))
	fmt.Println(strings.Index(test, "р"))
	fmt.Println(strings.Join([]string{test, s}, " - "))
	fmt.Println(strings.Repeat(test, 5))
	fmt.Println(strings.Replace(test, "ест", "***", -1))
	fmt.Println(strings.Split(test, "е"))
	fmt.Println(strings.ToUpper(test))
	fmt.Println(strings.Trim(test, "левоиТ")) // тримает края по любому символу последовательность не важна
	fmt.Println(strings.Trim("tetstet", "et"))
	fmt.Println(len(test), " байт в строке")
	fmt.Println(utf8.RuneCountInString(test), " символов в строке")

}
