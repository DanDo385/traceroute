// Package utils provides network utility functions for the traceroute application.
// This package handles DNS resolution, converting human-readable hostnames
// (like "google.com") into IP addresses that can be used for network communication.
package utils

import (
	"net"
)

// Resolve resolves a hostname to an IPv4 address.
// It performs a DNS lookup and returns the first IPv4 address found.
// If no IPv4 address is found, it returns an error.
func Resolve(host string) (string, error) {
	// Perform DNS lookup - this queries DNS servers to resolve the hostname
	// Returns all IP addresses (both IPv4 and IPv6) associated with the hostname
	ips, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}
	
	// Iterate through all returned IP addresses to find the first IPv4 address
	// IPv4 addresses are 32-bit addresses (e.g., 192.168.1.1)
	// IPv6 addresses are 128-bit addresses (e.g., 2001:0db8::1)
	// To4() converts an IP to IPv4 format, returning nil if it's IPv6
	for _, ip := range ips {
		if ip.To4() != nil {
			// Return the IPv4 address as a string (e.g., "192.168.1.1")
			return ip.String(), nil
		}
	}
	
	// If we reach here, no IPv4 address was found
	return "", &net.DNSError{Err: "no A records found", Name: host}
}
