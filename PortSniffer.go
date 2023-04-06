package main

import (
	"fmt"     //formatting input en output
	"net"     //word gebruikt om poorten te scannen
	"strconv" //converts van en naar strings van basis datatypes
	"time"    //word gebruikt om tijd bij te houden en weer te geven.
)

func main() {
	var startPort, endPort int

	fmt.Print("Enter IP address to scan: ")
	var ip string
	fmt.Scanln(&ip)

	fmt.Print("Enter start port: ")
	fmt.Scanln(&startPort)

	fmt.Print("Enter end port: ")
	fmt.Scanln(&endPort)

	scanPorts(ip, startPort, endPort, 500*time.Millisecond)
}

func scanPorts(ip string, startPort, endPort int, timeout time.Duration) {
	for port := startPort; port <= endPort; port++ {
		target := fmt.Sprintf("%s:%s", ip, strconv.Itoa(port))
		conn, err := net.DialTimeout("tcp", target, timeout)

		if err == nil {
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		} else {
			fmt.Printf("Port %d is closed\n", port)
		}
	}
}

//struct, array, flags
