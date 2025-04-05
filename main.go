package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <host> <port-range>")
		fmt.Println("Example: go run main.go scanme.nmap.org 20-100")
		return
	}

	host := os.Args[1]
	portRange := os.Args[2]

	start, end := parsePortRange(portRange)
	if start == -1 || end == -1 {
		fmt.Println("Invalid port range. Use format like 20-80.")
		return
	}

	fmt.Printf("🔍 Scanning %s from port %d to %d...\n\n", host, start, end)

	var wg sync.WaitGroup
	for port := start; port <= end; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", host, p)
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("✅ Port %d is open\n", p)
		}(port)
	}
	wg.Wait()
}

func parsePortRange(r string) (int, int) {
	parts := []rune(r)
	var startStr, endStr string
	seenDash := false

	for _, ch := range parts {
		if ch == '-' {
			seenDash = true
			continue
		}
		if !seenDash {
			startStr += string(ch)
		} else {
			endStr += string(ch)
		}
	}

	start, err1 := strconv.Atoi(startStr)
	end, err2 := strconv.Atoi(endStr)
	if err1 != nil || err2 != nil || start < 0 || end < start {
		return -1, -1
	}
	return start, end
}
