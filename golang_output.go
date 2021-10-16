package main
import "fmt"
func main(){
	var a rune = 'Ы'
	var b bool = false
	var i int = 11
	var f float32 = 11.22
	var a1 string = "123"
	var a2 string = "1234"
	fmt.Printf("%q", a) // вывод в кавычках
	fmt.Print("\n")
	fmt.Printf("%t", b) // вывод в булевых
	fmt.Print("\n")
	fmt.Printf("%b", i) // вывод целых в двоичной
	fmt.Print("\n")
	fmt.Printf("%d", i) // вывод целых
	fmt.Print("\n")
	fmt.Printf("%f", f) // вывод чисел с плавающей
	fmt.Print("\n")
	fmt.Printf("%.1f", f) // вывод чисел с плавающей
	fmt.Print("\n")
	fmt.Printf("%T - %T", b, f) // вывод типа
	fmt.Print("\n")
	fmt.Printf("%v - %v", b, f) // универсальный спецификатор
	fmt.Print("\n")
	fmt.Printf("%q \n%s", a1, a2)
	fmt.Print("\n")

	var sp float64 = 100.123456789
	result := fmt.Sprintf("%.2f", sp) //sprintf делает форматирование но не выводит результат
	fmt.Printf("%q\n", result) // вывод: "100.12"
}