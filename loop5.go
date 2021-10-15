/*
Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
*/
package main
import "fmt"
func main(){
	var i,n,c,d uint

	fmt.Scan(&n, &c, &d)

	for i = 1; i <= n; i++ {
		if i % c == 0 && i % d != 0 {
			fmt.Println(i)
			break
		}
	}

}