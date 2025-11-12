package main

import (
	"fmt"
	"log"
	"os"

	"traceroute/internal/netutil"
	"traceroute/internal/tracer"
)

// main is the entry point of the traceroute application.
// Traceroute is a network diagnostic tool that traces the path packets take
// from your computer to a destination host, showing each intermediate router
// (hop) along the way.
func main() {
	// Check if a hostname argument was provided
	// os.Args[0] is the program name, os.Args[1] would be the first argument
	if len(os.Args) < 2 {
		log.Fatalf("Usage: $s <host>", os.Args[0])

	}

	// Extract the hostname from command-line arguments
	// Example: "traceroute google.com" -> host = "google.com"
	host := os.Args[1]
	fmt.Printf("Traceroute to %s\n", host)
	// Resolve the hostname to an IP address using DNS lookup
	// DNS (Domain Name System) translates human-readable domain names
	// (like "google.com") into IP addresses (like "142.250.191.14")
	// that computers use to route packets across the internet
	ip, err := netutil.ResolveHost(host)
	if err != nil {
		log.Fatalf("Failed to resolve host: %v", err)
	}

	// Display the resolved IP address
	fmt.Printf("Resolved %s -> %s\n", host, ip)

	// Run the traceroute
	tracer.Run(ip)
}