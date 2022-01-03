package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"log"
)

func main() {
	rc, err := zip.OpenReader("files/task.zip")
	if err != nil {
		log.Fatal(err)
	}

	defer rc.Close()
	//fmt.Println(rc.File[10].Name)
	//fmt.Println(rc.File[10].FileInfo().Size())
	for _, f := range rc.File {
		if f.FileInfo().IsDir() == false {
			fs, _ := f.Open()

			a, err := csv.NewReader(fs).ReadAll()
			if err == nil && len(a) > 1 {
				fmt.Println(f.Name)
				fmt.Println(a[4][2])
			}

		}
	}
}
