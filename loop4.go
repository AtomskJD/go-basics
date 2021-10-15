/*
Последовательность состоит из натуральных чисел и завершается числом 0. 
Определите количество элементов этой последовательности, которые равны ее 
наибольшему элементу.
*/
package main
import "fmt"
func main(){
	var n, max, len int

		
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max {
			max = n
			len = 0
		} else if n == max {
			len++
		} 
	}

	fmt.Println(len + 1)

}