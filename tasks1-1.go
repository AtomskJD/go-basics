package main
import "fmt"
func main(){
	var a, ai, b int
	fmt.Scan(&a)
	for i := 0; i < 3; i++ {
		ai = a % 10
		a = a / 10
		b = b + ai 
	}
	fmt.Println(b)
}