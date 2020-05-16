package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := os.Args[1]
	fileinfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	if fileinfo.IsDir() {
		list, _ := ioutil.ReadDir(path)
		for _, item := range list {
			fmt.Printf("dir: %s\n", item.Name())
		}
	}
	fmt.Println("корневая папка ", fileinfo.Name())

}
