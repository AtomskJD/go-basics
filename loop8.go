/*
Даны два числа. Определить цифры, входящие в запись как первого, 
так и второго числа.
*/
package main
import "fmt"
func main(){
  var a, b, bsave, da, db, result, dresult uint
  fmt.Scan(&a, &b)
  bsave = b
  for a > 0 {
    da = a % 10
    a = a / 10
    b = bsave
    for b > 0 {
      db = b % 10
      b = b / 10
      if da == db {
        result = result * 10 + db
      }
    }
  }

  for result > 0 {
    dresult = result % 10
    result = result / 10
    fmt.Print(dresult)
    fmt.Print(" ")
  }
}