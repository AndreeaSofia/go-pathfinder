package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		ShowUsage()
		return
	}

	host := os.Args[1]
	portRange := os.Args[2]

	start, end, err := ParsePortRange(portRange)
	if err != nil {
		ExitWithError("Invalid port range. Use format like 20-80.")
	}

	fmt.Printf("🔍 Scanning %s from port %d to %d...\n\n", host, start, end)
	ScanPorts(host, start, end)
}
