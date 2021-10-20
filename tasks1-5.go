package main
import "fmt"
func main(){
	var a, b, c uint
	fmt.Scan(&a, &b, &c)
	if a + b > c {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}

}