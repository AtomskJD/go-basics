package main
import "fmt"
func main(){
	var t uint
	fmt.Scan(&t)
	if t > 0 && t <= 86400 {
		fmt.Printf("It is %d hours %d minutes.", t / 60 / 60, (t/60) - (t/60/60 * 60) )
	}
}