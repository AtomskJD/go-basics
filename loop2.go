package main
import "fmt"
func main(){
	var a, b, sum uint
	fmt.Scan(&a, &b)
	if a < b && a <= 100 && b <= 100 {
		for ; a <= b; a++ {
			sum = sum + a
		}
	}
	fmt.Println(sum)
}