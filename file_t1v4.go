package main

import (
	"archive/zip"
	"bufio"
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

			sc := bufio.NewScanner(fs)

			for sc.Scan() {
				if []byte(sc.Text())[0] != 0 {
					fmt.Printf("%v\n", sc.Text())
				}
			}

		}
	}
}
