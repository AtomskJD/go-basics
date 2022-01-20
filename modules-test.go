package main

import (
	"atomskjd/gomodule"
	. "fmt"
	"github.com/atomskjd/modules-test/bye"
	"github.com/atomskjd/modules-test/hello"
)

func main() {
	gomodule.Test()
	Println("HELLO")
	Println(hello.Jp())
	Println(bye.Jp())
}
