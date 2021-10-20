/*
  Второй массив содержит пары индексов для свапа в первом массиве
*/

package main
import "fmt"
func main(){
  var a []int
  fmt.Println(a)
  a = []int{3:11}
  fmt.Println(a)
  a[3] =4
  fmt.Println(a)
  fmt.Printf("%T", a)

  // make делает срез
  a = make([]int, 5, 10) // [0 0 0 0 0 0 0 0 0 0]
  a[4] = 123
  fmt.Println(a)

  var initialUsers = [8]string{"Катя","Маша","Вера","ВИка","Марина","Артур","Илья","Денис"}
  users1 := initialUsers[2:4]
  users2 := initialUsers[:3]
  users3 := initialUsers[3:]
  users5 := initialUsers[4:5]


  // users4 := initialUsers[:]

  fmt.Println(initialUsers)
  fmt.Println(users1)
  fmt.Println(users2)
  users2 = append(users2, "Сара", "Влад")
  fmt.Println(users2)

  fmt.Println(users3)
  users5 = append(users5, "Игорь", "Ната")
  fmt.Println(cap(users5), users5)
  fmt.Println(initialUsers)
}