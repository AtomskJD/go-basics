package main

import "fmt"

type battery struct {
	BinnaryCharge []byte
}

func (b battery) String() string {
	batteryTemplate := "XXXXXXXXXX"
	r := []rune(batteryTemplate)
	rs := []rune(" ")
	bi := uint8(0)
	for i := 0; i < 10; i++ {
		if b.BinnaryCharge[i] == 48 {
			r[bi] = rs[0]
			bi++
		}
	}
	return fmt.Sprintf("[%v]", string(r))
}

func main() {
	var test string
	fmt.Scan(&test)
	bt := []byte(test)
	batteryForTest := battery{bt}
	fmt.Println(batteryForTest)

}
