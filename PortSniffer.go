package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// PortConfig represents the start and end port range to scan
type PortConfig struct {
	StartPort int `json:"StartPort"`
	EndPort   int `json:"EndPort"`
}

func main() {
	// Read the port configuration from the Portconfig.json file
	portConfig, err := readPortConfig("Portconfig.json")
	if err != nil {
		fmt.Println("Error reading Portconfig.json file:", err)
		return
	}

	// Get the IP address to scan from the user
	ip, err := getUserInput()
	if err != nil {
		fmt.Println("Error reading user input:", err)
		return
	}

	// Scan the specified ports for the given IP address
	openPorts := scanPorts(ip, portConfig.StartPort, portConfig.EndPort, 500*time.Millisecond)

	// If no open ports were found, print a message and exit
	if len(openPorts) == 0 {
		fmt.Println("No open ports found.")
		return
	}

	// Print each open port that was found
	for _, port := range openPorts {
		fmt.Printf("Port %d is open\n", port)
	}
}

// readPortConfig reads the port configuration from the specified JSON file
func readPortConfig(filename string) (*PortConfig, error) {
	portConfig := &PortConfig{}

	// Open the specified file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON data in the file into a PortConfig struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(portConfig)
	if err != nil {
		return nil, err
	}

	return portConfig, nil
}

// getUserInput prompts the user to enter an IP address to scan
func getUserInput() (string, error) {
	var ip string
	fmt.Print("Enter IP address to scan: ")
	_, err := fmt.Scanln(&ip)
	if err != nil {
		return "", err
	}
	return ip, nil
}

// scanPort scans the specified port for the given IP address and returns true if it is open
func scanPort(ip string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}

// scanPorts scans the specified port range for the given IP address and returns a slice of open port numbers
func scanPorts(ip string, startPort, endPort int, timeout time.Duration) []int {
	var wg sync.WaitGroup
	openPorts := []int{}

	// Scan each port in the specified range in a separate goroutine
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			if scanPort(ip, port, timeout) {
				openPorts = append(openPorts, port)
			}
		}(port)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return openPorts
}

//tekst edit kijken of het werkt
