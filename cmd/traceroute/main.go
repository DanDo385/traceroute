package main

import (
	"fmt"
	"os"

	"traceroute/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: traceroute <hostname>")
		return
	}

	host := os.Args[1]

	ip, err := utils.Resolve(host)
	if err != nil {
		fmt.Printf("DNS lookup failed: %v\n", err)
		return
	}

	fmt.Printf("Resolved %s -> %s\n", host, ip)
}
