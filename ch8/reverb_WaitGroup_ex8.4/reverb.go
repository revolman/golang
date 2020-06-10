package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
	// тут должен быь wg.Done(), но его заменит анонимная функция
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func(shout string) {
			defer wg.Done()
			echo(c, shout, 1*time.Second)
		}(input.Text()) // обязательно нужно передавать текст параметром функции
	}

	go func() {
		wg.Wait()
		c.(*net.TCPConn).CloseWrite()
	}()
	// c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}
