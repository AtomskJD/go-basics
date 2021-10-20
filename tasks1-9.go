package main
import "fmt"
func main(){
	var a,b  uint32
	fmt.Scan(&a)
	for {
		b += a % 10
		a = a / 10

		if a < 1 {
			if b > 10 {
				a = b
				b = 0
			} else {
				break
			}
		}
	}

	fmt.Println(b)
}