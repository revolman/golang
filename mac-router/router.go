package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	iface, err := getIface()
	if err != nil {
		log.Fatal(err)
	}

	var hosts = []string{"192.168.0.10", "192.168.0.160", "192.168.0.222"}

	for _, host := range hosts {
		cmd := exec.Command("route", "add", "-host", host, "-interface", iface)
		out, err := cmd.CombinedOutput()

		if err != nil {
			log.Fatalf("Не удалось выполнить системный вызов: %v", err)
		}

		fmt.Printf("%s", string(out))
	}
}

func getIface() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, iface := range ifaces {

		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatal(err)
		}

		for _, addr := range addrs {
			addrStr := addr.String()
			split := strings.Split(addrStr, "/")
			ip := net.ParseIP(split[0])

			if strings.Contains(ip.String(), "10.0.8.") { // 10.0.8. для OpenVPN
				// fmt.Printf("%v\t %s\n", iface.Name, ip)	// debug
				return iface.Name, nil
			}
		}
	}
	return "", fmt.Errorf("Не получилось обнаружить VPN-интерфейс")
}
