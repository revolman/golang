// Упражнение 8.1 - Clockwall. Возможно не самая оптимизованная реализация,
// зато точно соответствует условиям задачи.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// ZoneInfo инфа о тайм зоне
type ZoneInfo struct {
	Name        string // название
	Address     string // сетевой адрес
	CurrentTime string // текущее время
}

func main() {
	var timeZones []*ZoneInfo // все полученные зоны

	if len(os.Args) < 1 {
		usage()
	}
	// подготавливаю список полученных зон, переданных в аргументе
	args := os.Args[1:]
	for _, arg := range args {
		tz := &ZoneInfo{}
		if !strings.Contains(arg, "=") {
			usage()
		}
		seq := strings.Split(arg, "=")
		tz.Name = seq[0]
		tz.Address = seq[1]
		timeZones = append(timeZones, tz)
	}

	clockWall(timeZones)
}

// в качестве решения нужно попробовать обновлять время в структуре для каждого
// отдельного *ZoneInfo, а результат выводить в бесконечном цикле не давай главной подпрограмме завершиться
func clockWall(tz []*ZoneInfo) {
	for _, zone := range tz {
		conn, err := net.Dial("tcp", zone.Address)
		if err != nil {
			log.Fatal(err)
		}
		go refreshTime(conn, zone)
	}

	for {
		for _, zone := range tz {
			fmt.Printf("%s: %s\t", zone.Name, zone.CurrentTime)
		}
		fmt.Print("\r")
		time.Sleep(200 * time.Millisecond)
	}
}

// подпрограмма обновления текущего времени
func refreshTime(conn net.Conn, zone *ZoneInfo) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for {
		for scanner.Scan() {
			zone.CurrentTime = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func usage() {
	fmt.Printf("Введите аргументы:\nname=address:port\n")
	os.Exit(1)
}
