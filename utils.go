package main

import (
	"fmt"
	"os"
)

func ShowUsage() {
	fmt.Println("Usage: go run main.go <host> <port-range>")
	fmt.Println("Example: go run main.go scanme.nmap.org 20-100")
}

func ExitWithError(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	os.Exit(1)
}
