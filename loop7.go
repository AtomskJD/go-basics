/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов,
 после чего дробная часть копеек отбрасывается. Каждый год сумма вклада 
 становится больше. Определите, через сколько лет вклад составит не менее y 
 рублей.


*/
package main
import "fmt"
func main(){
  var x,p,y float32
  var i uint
  fmt.Scan(&x)
  fmt.Scan(&p)
  fmt.Scan(&y)
  pp := p / 100
  for x < y {
    i++
    x = x * pp + x
    // if x >= y {
    //   break
    // }
  }
  fmt.Println(i)
}