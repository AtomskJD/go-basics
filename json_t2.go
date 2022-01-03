package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Root []struct {
	GlobalID int64 `json:"global_id"`
}

func main() {
	var root Root
	var sum int64
	file, _ := os.Open("files/data-20190514T0100.json")
	in, err := ioutil.ReadAll(file)
	if err == nil {
		if json.Valid(in) {
			fmt.Println("VALID")
			json.Unmarshal(in, &root)

			for _, st := range root {
				sum += st.GlobalID
			}
			fmt.Println(sum)
		}
	}
}
