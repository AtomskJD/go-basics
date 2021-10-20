package main
import "fmt"
func main(){
  var a [3]int = [3]int{1,2,3}
  b := [3]int{1,2,3}
  c := [...]int{2,3,4}
  d := [3]int{2: 12}

  fmt.Println(a, b, c, d)
  // сравнение массивов
  fmt.Println(a == b)
  fmt.Println(a == c)

  e := [5]int{1, 2, 3, 4, 5}
  fmt.Println(e)

  for i := 0; i < len(e); i++ {
    fmt.Println(e[i])
  }

  fmt.Println(e)

  for idx, elem := range e {
      fmt.Printf("Элемент с индексом %d: %d\n", idx, elem)
  }
}