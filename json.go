package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Students
}
type Students struct {
	LastName  string
	FirstName string
	Rating    []int
}
type Rating struct {
	Average float64
}

func main() {
	var group Group
	//file, _ := os.Open("files/data.json")
	in, err := ioutil.ReadAll(os.Stdin)
	if err == nil {
		if json.Valid(in) {
			json.Unmarshal(in, &group)
		}
		countStudents := len(group.Students)
		lenRatings := 0
		for _, st := range group.Students {
			lenRatings += len(st.Rating)
		}

		average := fmt.Sprintf("%0.1f", float32(lenRatings)/float32(countStudents))
		averageFl, _ := strconv.ParseFloat(average, 64)
		res := Rating{Average: averageFl}

		data, err := json.MarshalIndent(res, "", "\t")
		if err != nil {
			fmt.Println(err)
			return
		}

		os.Stdout.Write(data)
	}
}
