/*Напишите программу, принимающая на вход число N(N≥4), 
а затем N целых чисел-элементов среза. 
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.*/
package main
import "fmt"
func main(){
	var N, ai, i int
	var a []int

	fmt.Scan(&N)
	if N >= 4 {
		for i = 0; i < N; i++ {
			fmt.Scan(&ai)
			a = append(a, ai)
		}

		fmt.Println(a[3])
	}
}