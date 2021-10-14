package main
import "fmt"
func main(){
	var a uint
	fmt.Scan(&a)
	if a >= 10000 {
		fmt.Println(a / 10000)
	} else if a >= 1000 {
		fmt.Println(a / 1000)
	} else if a >= 100 {
		fmt.Println(a / 100)
	} else if a >= 10 {
		fmt.Println(a / 10)
	} else {
		fmt.Println(a)
	}
}