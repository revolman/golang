package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

// ZoneInfo имя таймзоны и адресс сервера
type ZoneInfo struct {
	Name    string
	Address string
}

const format = ("%v\t%v\n")

func main() {
	var timeZones []*ZoneInfo // все полученные зоны

	if len(os.Args) < 1 {
		fmt.Printf("Введите аргументы:\nname=address:port\n")
		os.Exit(1)
	}
	args := os.Args[1:]
	for _, arg := range args {
		tz := &ZoneInfo{}
		seq := strings.Split(arg, "=")
		tz.Name = seq[0]
		tz.Address = seq[1]
		timeZones = append(timeZones, tz)
	}

	clockWall(timeZones)
}

func clockWall(tz []*ZoneInfo) {
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Printf(format, "Зона", "Время")
	fmt.Printf(format, "----", "-----")

	for _, zone := range tz {
		// fmt.Fprintf(tw, format, zone.Name, go refreshTime(zone))
		conn, err := net.Dial("tcp", zone.Address)
		if err != nil {
			log.Fatal(err)
		}

		go timePrint(conn, zone)

		// go refreshTime(tw, conn, zone)
	}
	tw.Flush()

	time.Sleep(5 * time.Minute)
}

// в качестве решения нужно попробовать обновлять время в структуре для каждого
// отдельного *ZoneInfo, а в главной функции выводить результат в бесконечном цикле
func timePrint(conn net.Conn, zone *ZoneInfo) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for {
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Fprintf(os.Stdout, "\r%s: %s", zone.Name, text)
		}
	}
}

// func refreshTime(tw *tabwriter.Writer, conn net.Conn, zone *ZoneInfo) {
// 	defer conn.Close()

// 	var tm = &bytes.Buffer{}

// 	if _, err := io.Copy(tm, conn); err != nil {
// 		log.Fatal(err)
// 	}
// 	newTime, err := tm.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Fprintf(tw, "\r%v\t%v", zone.Name, newTime)

// fmt.Fprintf(tw, "\r%v\t%s", zone.Name, conn)
// time.Sleep(1 * time.Second)

// }
