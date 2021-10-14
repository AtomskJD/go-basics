package main
import "fmt"
func main(){
	i := 0
	switch i {
	case 0:
		fmt.Println("zero")
		fallthrough
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("default")
		return
	}


	var c uint32
	fmt.Scan(&c)
	switch {
	case 1 <= c && c <= 9:
		fmt.Println("от 1 до 9")
	case 10 <= c && c <= 250:
		fmt.Println("от 10 до 250")
	case 1000 <= c && c <= 6000:
		fmt.Println("от 1000 до 6000")
	}
}