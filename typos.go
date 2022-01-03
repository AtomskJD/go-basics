package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxInt32)
	a := 5.0 / 2
	fmt.Println(a) // 2
	b := 101.0 / 2.0
	fmt.Println(b + float64(uint8(b)))

	a1 := "12"
	//c1 := 100
	b1 := []byte(a1)
	a1 += string(b1)
	fmt.Println(a1)
}
func convert(x int64) uint16 {
	return uint16(x)
}
