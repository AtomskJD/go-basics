/**
Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
и возвращающую количество полученных функцией аргументов и их сумму.
*/
package main

//func main() {
//	a, b := sumInt(1, 0, 15, 55, 22, 4, 7)
//	fmt.Println(a, b)
//}

func sumInt(i ...int) (int, int) {
	var sum int = 0
	for _, el := range i {
		sum += el
	}
	return len(i), sum
}
