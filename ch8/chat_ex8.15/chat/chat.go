// Большую часть слямзил тут: https://github.com/kdama/gopl/blob/master/ch08/ex15/main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const timeout = 5 * time.Minute
const outbuffer = 64 // Размер буффера канала сообщений для передачи

type client struct {
	name  string
	outch chan<- string // канал для отправки  | Writer
	inch  chan<- string // канал для получения | Reader
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // Все входящие сообщения клиента
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

func broadcaster() {
	clients := make(map[client]bool) // Все подключенные клиенты
	for {
		select {
		case msg := <-messages:
			// Широковещательное входящее сообщение во все
			// каналы исходящих сообщений для клиентов.
			for cli := range clients {
				select {
				case cli.outch <- msg:
				default:
					// пропустить, если сообщение долго не приходит
				}
			}
		case cli := <-entering:
			clients[cli] = true

			// Оповестить вошедшего о клиентах онлайн
			var online []string
			for client := range clients {
				online = append(online, client.name)
			}
			cli.outch <- fmt.Sprintf("%d клиентов: %s", len(online), strings.Join(online, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outch)
		}
	}
}

func handleConn(conn net.Conn) {
	inch := make(chan string)
	outch := make(chan string, outbuffer) // Исходящие сообщения клиентов

	go clientWriter(conn, outch)
	go clientReader(conn, inch)

	// Спрашиваем имя
	var who string

	outch <- "Введите ваше имя: "

	// Отключить клиента, который отказался представиться
	select {
	// за наполнение inch отвечает clientReader
	case in, ok := <-inch:
		if !ok {
			conn.Close()
			return
		}
		who = in
	// Отключение по таймауту
	case <-time.After(timeout):
		conn.Close()
		return
	}

	messages <- who + " подключился"
	entering <- client{who, outch, inch}

	// обработка ввода клиента, читает записи из inch
	for {
		select {
		case text, ok := <-inch:
			if ok {
				messages <- who + ": " + text
			} else {
				leaving <- client{who, outch, inch}
				messages <- who + " отключился"
				conn.Close()
				return
			}
		case <-time.After(timeout):
			leaving <- client{who, outch, inch}
			messages <- who + " отключился"
			conn.Close()
			return
		}
	}
}

//!+ А не переименовать ли эти функции? :)

// Записывает полученные сообщения в соединение
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // Игнорирование ошибок сети
	}
}

// "Читает" ввод в канал
func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
	close(ch)
	conn.Close()
}

//!-
