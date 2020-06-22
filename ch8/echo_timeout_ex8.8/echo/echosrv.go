// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Reverb1 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//!+
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func scan(c net.Conn, out chan<- string) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		out <- input.Text()
	}
	close(out)
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go func(c net.Conn) {
			var shout = make(chan string, 1)
			timer := time.After(10 * time.Second)

			go scan(c, shout)

			for done := false; !done; {
				select {
				case text := <-shout:
					go echo(c, text, 1*time.Second)
					timer = time.After(10 * time.Second)
				case <-timer:
					done = true
				}
			}
			// ticker.Stop()
			fmt.Fprintln(c, "write closed: timeout")
			c.(*net.TCPConn).CloseWrite()

			// for input.Scan() {
			// 	go echo(c, input.Text(), 1*time.Second)
			// }
			// // NOTE: ignoring potential errors from input.Err()
			// c.Close()
		}(conn)
	}
}
