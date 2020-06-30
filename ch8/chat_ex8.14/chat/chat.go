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

type client struct {
	name string
	ch   chan<- string
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
				cli.ch <- msg
			}
		case cli := <-entering: // управление переменной передаётся бродкастеру, что бы избежать состояния гонки
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
	check := make(chan struct{}) // канал создан для проверки было ли отправлено сообщение

	ch := make(chan string) // Исходящие сообщения клиентов
	go clientWriter(conn, ch)
	input := bufio.NewScanner(conn)

	// Спрашиваем имя
	var who string
	go func() {
		ch <- "Введите ваше имя: "
		if input.Scan() {
			who = input.Text()
			check <- struct{}{}
		} else {
			leaving <- client{who, ch}
			messages <- who + " отключился"
			conn.Close()
			return
		}
	}()

	// Отключить клиента, который отказался представиться
loop:
	for {
		select {
		case _, ok := <-check:
			if ok {
				break loop // разорвать loop, просто breake разорвёт select
			} else {
				conn.Close()
				return
			}
		case <-time.After(timeout):
			conn.Close()
			return
		}
	}

	messages <- who + " подключился"
	entering <- client{who, ch}

	go func() { // отдельная подпрограмма, которая обрабатывает ввод клиента, необходимо выполнять параллельно
		for {
			if input.Scan() {
				messages <- who + ": " + input.Text()
				check <- struct{}{} // это маркер для того что сообщение было отправлено
			} else {
				// Игнорируем потенциальные ошибки input.Err()
				leaving <- client{who, ch}
				messages <- who + " отключился"
				conn.Close()
				return
			}
		}
	}()

	for {
		select {
		case _, ok := <-check: // если сообщение было отправлено, тогда ждать следующего
			if !ok { // дополнительная, но не обязательная проверка отправлено ли сообщение
				conn.Close()
				return
			}
		case <-time.After(timeout): // если в течении 10 сек не было собщения, time.After получит значение
			conn.Close() // и закроет соединение
			return
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // Игнорирование ошибок сети
	}
}
