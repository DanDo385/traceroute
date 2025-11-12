package main

import (
	"fmt"
	"os"

	"traceroute/internal/utils"
)

// main is the entry point of the traceroute application.
// Traceroute is a network diagnostic tool that traces the path packets take
// from your computer to a destination host, showing each intermediate router
// (hop) along the way.
func main() {
	// Check if a hostname argument was provided
	// os.Args[0] is the program name, os.Args[1] would be the first argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: traceroute <hostname>")
		return
	}

	// Extract the hostname from command-line arguments
	// Example: "traceroute google.com" -> host = "google.com"
	host := os.Args[1]

	// Resolve the hostname to an IP address using DNS lookup
	// DNS (Domain Name System) translates human-readable domain names
	// (like "google.com") into IP addresses (like "142.250.191.14")
	// that computers use to route packets across the internet
	ip, err := utils.Resolve(host)
	if err != nil {
		fmt.Printf("DNS lookup failed: %v\n", err)
		return
	}

	// Display the resolved IP address
	fmt.Printf("Resolved %s -> %s\n", host, ip)
}
