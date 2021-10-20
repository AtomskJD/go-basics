package main
import "fmt"
func main(){
	var a,b int32
  var f bool = false
	
	fmt.Scan(&a, &b)
  if a <= b {
			for ; a <= b; b-- {
				if b % 7 == 0 {
					fmt.Println(b)
		            f = true
		            break
				} 
			}
		if f == false {
		    fmt.Println("NO")
	  }
  }  
}