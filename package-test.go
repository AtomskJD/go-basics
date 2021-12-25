package main

import (
	"atomskjd/testmodule"
	. "fmt"
)

func main() {
	testmodule.Gotest()
	a, b := sumInt(2, 2, 5)
	Println(a, b)
	Println("HELLO")
}
