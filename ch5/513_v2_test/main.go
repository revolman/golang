package main

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

func main() {
	query := "http://golang.org/doc/dsa/"
	url, _ := url.Parse(query)

	var dir string
	dir = url.Host
	var filename string

	fmt.Println(url.Path, url.Host, url.RawPath)

	if strings.HasSuffix(url.Path, ".html") {
		parts := strings.Split(url.Path, "/")
		fmt.Println(parts[len(parts)-1])
		dir = dir + strings.Join(parts[:len(parts)-1], "/")
		fmt.Println(dir)
	} else {
		dir = dir + url.Path
		filename = "index.html"
	}
	fpath := filepath.Join(dir, filename)
	fmt.Println(fpath)
}
