package main
import "fmt"
func main(){
	var a, b, c, d int
	fmt.Scan(&a)

	b = a / 100 % 10
	c = a /10 % 10
	d = a % 10

	switch {
	case b == c:
		fallthrough
	case b == d:
		fallthrough
	case d == c:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
		return
	}
}