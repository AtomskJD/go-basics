package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
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

			tf, _ := ioutil.ReadAll(fs)
			if tf[0] != 0 {
				fmt.Printf("%v %d\n", f.Name, f.FileInfo().Size())
				fmt.Printf("%s\n", tf)
			}
		}
	}
}
