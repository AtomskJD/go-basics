package main

import "fmt"

func main() {
	var value1, value2, operation interface{}
	value1 = true
	value2 = 3.6
	operation = "/"

	if testVar(value1) == true && testVar(value2) == true && testOperator(operation) == true {
		switch operation {
		case "+":
			res := value1.(float64) + value2.(float64)
			fmt.Printf("%.4f", res)
			break
		case "-":
			res := value1.(float64) - value2.(float64)
			fmt.Printf("%.4f", res)
			break
		case "*":
			res := value1.(float64) * value2.(float64)
			fmt.Printf("%.4f", res)
			break
		case "/":
			res := value1.(float64) / value2.(float64)
			fmt.Printf("%.4f", res)
			break
		}
	}

}

func testVar(i interface{}) bool {
	if _, ok := i.(float64); ok {
		return true
	}
	fmt.Printf("value=%v: %T", i, i)

	return false
}

func testOperator(i interface{}) bool {
	if v, ok := i.(string); ok {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			return true
		} else {
			fmt.Println("неизвестная операция")
			return false
		}
	}
}
