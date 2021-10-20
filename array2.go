/*
  Второй массив содержит пары индексов для свапа в первом массиве
*/

package main
import "fmt"
func main(){
  var workArray [10]byte
  var swa [6]byte
  var tmp byte


  for i:=0; i<len(workArray); i++ {
    fmt.Scan(&workArray[i])
  }
  for i:=0; i<len(swa); i++ {
    fmt.Scan(&swa[i])
  }
  for i:=0; i < len(swa); i = i + 2 {
    tmp           = workArray[swa[i]]
    workArray[swa[i]]   = workArray[swa[i+1]]
    workArray[swa[i+1]] = tmp

  }

  for _, elem := range workArray {
    fmt.Printf("%d ", elem)
  }
}