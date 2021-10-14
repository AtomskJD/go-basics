package main
import "fmt"
func main(){
  var a uint
  fmt.Scan(&a)
  l := a / 100000 + a / 10000 % 10 + a / 1000 % 10
  r := a % 1000 / 100 + a % 100 / 10 + a % 10
  fmt.Println(l, r)
  if l == r {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }
}