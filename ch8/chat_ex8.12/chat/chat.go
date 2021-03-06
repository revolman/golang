package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name string
	ch   chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // Все входящие сообщения клиента
)

func broadcaster() {
	clients := make(map[client]bool) // Все подключенные клиенты
	for {
		select {
		case msg := <-messages:
			// Широковещательное входящее сообщение во все
			// каналы исходящих сообщений для клиентов.
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true

			// Оповестить вошедшего о клиентаъ онлайн
			var online []string
			for client := range clients {
				online = append(online, client.name)
			}
			cli.ch <- fmt.Sprintf("%d клиентов: %s", len(online), strings.Join(online, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // Исходящие сообщения клиентов
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()

	ch <- "Вы " + who
	messages <- who + " подключился"
	entering <- client{who, ch}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// Игнорируем потенциальные ошибки input.Err()

	leaving <- client{who, ch}
	messages <- who + " отключился"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // Игнорирование ошибок сети
	}

}
