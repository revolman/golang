package main

import (
	"fmt"
	"sync"
)

var cache struct {
	sync.Mutex
	mapping map[string]string
}

func lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	cache.mapping = make(map[string]string)
	cache.mapping["hello"] = "hi"
	v := lookup("hello")
	fmt.Println(v)
}
