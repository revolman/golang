// Упражнение 8.3
package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	// можно сразу предопределить тип соединения,
	// тогда методы будут джоступны без декларации.
	// addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")
	// conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn) // игнор ошибок
		log.Panicln("Done")
		done <- struct{}{} // сигнал главной  подпрограмме
	}()

	mustCopy(conn, os.Stdin)
	//!+test
	time.Sleep(5 * time.Second)
	os.Stdin.Close()
	//!-test
	// можно использовать декларацию типов для проверки типа соединения,
	// и если тип *net.TCPConn, то можно использовать его методы
	conn.(*net.TCPConn).CloseWrite()
	<-done // Ожидание завершения фоновой go-подпрограммы
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
