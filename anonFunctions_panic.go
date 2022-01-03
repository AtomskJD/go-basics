package main

import (
	"errors"
	"fmt"
)

func main() {
	ExamplePanicRecover()
}

func someFuncWithPanic() (err error) {
	defer func() {
		// отложенный вызов анонимной функции, проверяющей, что работа функции завершена
		// без ошибок. Если функция recover() возвращает что угодно кроме nil, значит в ходе
		// выполнения функции возникла паника.
		if e := recover(); e != nil {
			// Здесь происходит приведение интерфейса (об этом мы расскажем буквально в
			// следующем уроке. Результат приведения присваивается переменной err типа error
			// которая уже объявлена при самом вызове функции someFuncWithPanic.
			err = e.(error)

			// после этого анонимная функция завершает свою работу, паника обработана,
			// переменная err, в которой содержится информации о возникшей панике,
			// возвращается как результат выполнения функции.
		}
	}()

	panic(errors.New("fatal error"))
	return
}

func ExamplePanicRecover() {
	if err := someFuncWithPanic(); err != nil {
		fmt.Println(err)
	}

	// Output:
	// fatal error
}
