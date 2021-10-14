package main
import "fmt"
func main(){
	a := 5
	b := 5
	if a < b {
		fmt.Println("а меньше б")
	} else if a > b {
		fmt.Println("б меньше а")
	} else {
		fmt.Println("они равны")
	}
}