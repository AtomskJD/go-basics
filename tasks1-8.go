package main
import "fmt"
func main(){

		array := []int{}
		var a, n, count int
		fmt.Scan(&n)
		for i:=0; i < n; i++{
			fmt.Scan(&a)
			array = append(array, a)
		}

		var min int
		min = array[0]
		for _, elem := range array {
			if elem < min {
				min = elem
				count = 1
			} else if elem == min {
				count++
			}
		}

		fmt.Println(count)
}