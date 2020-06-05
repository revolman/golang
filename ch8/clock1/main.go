package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // этот объект прослушивает 8000 порт tcp
	if err != nil {
		log.Fatal(err)
	}
	// Бесконечный цикл ожидает соединения
	for {
		conn, err := listener.Accept() // Accept блокируется до тех пор, пока не будет установено соединение
		if err != nil {
			log.Print(err)
			continue
		}
		// handleConn(conn) // Обработка едиственного соединения
		go handleConn(conn) // Обработка каждого соединения в отдельно подпрограмме
	}
}

// net.Conn соответствует интерфейсу io.Writer, а значит можно осуществлять вывод прямо в него.
func handleConn(c net.Conn) {
	defer c.Close() // отложенное закрытие соединения
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil { // если не удалось выполнить запись
			return // Например, отключение клиента
		}
		time.Sleep(1 * time.Second) // ожидать 1 секунду
	}
}
