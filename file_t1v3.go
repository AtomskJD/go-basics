package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
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

			rd := bufio.NewReader(fs)
			for s, err := rd.ReadString('\n'); err == nil && err != io.EOF; {

				fmt.Printf("%s\n", s)
				s, err = rd.ReadString('\n')
			}
		}
	}
}
