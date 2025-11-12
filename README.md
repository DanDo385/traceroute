# Traceroute - Network Path Tracing Tool

A network diagnostic tool written in Go that traces the path packets take from your computer to a destination host on the internet. This project demonstrates fundamental networking concepts including DNS resolution, IP addressing, and packet routing.

## Table of Contents

- [What is Traceroute?](#what-is-traceroute)
- [How Traceroute Works](#how-traceroute-works)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [How It Works (Technical Details)](#how-it-works-technical-details)
- [Building from Source](#building-from-source)
- [Development](#development)

## What is Traceroute?

**Traceroute** is a network diagnostic tool that reveals the path packets take as they travel from your computer to a destination host on the internet. When you visit a website or connect to a server, your data doesn't travel directly—it hops through multiple routers (network devices) before reaching its destination.

### Why is Traceroute Useful?

- **Network Troubleshooting**: Identify where network problems occur along the path
- **Performance Analysis**: See how many hops your data takes and measure latency
- **Understanding Internet Routing**: Learn how the internet routes traffic
- **Security Analysis**: Discover the path your data takes through the network

### Real-World Example

When you visit `google.com` from your computer, your request might travel through:
1. Your home router (192.168.1.1)
2. Your ISP's router
3. Regional network routers
4. Google's edge servers
5. Finally, Google's data center

Traceroute shows you each of these "hops" and how long each one takes.

## How Traceroute Works

### The Core Concept: TTL (Time To Live)

Traceroute uses a clever technique involving the **TTL (Time To Live)** field in IP packets:

1. **TTL Field**: Every IP packet has a TTL field that decrements at each router hop
2. **When TTL Reaches Zero**: The router discards the packet and sends back an ICMP "Time Exceeded" message
3. **Incremental Probing**: Traceroute sends packets with TTL=1, then TTL=2, then TTL=3, etc.
4. **Mapping the Path**: Each router that responds reveals itself, building a map of the network path

### Step-by-Step Process

```
Your Computer                    Router 1              Router 2              Destination
     |                              |                      |                      |
     |--[TTL=1]-------------------->|                      |                      |
     |<--[ICMP Time Exceeded]-------|                      |                      |
     |                              |                      |                      |
     |--[TTL=2]------------------------------------------>|                      |
     |<--[ICMP Time Exceeded]------------------------------|                      |
     |                              |                      |                      |
     |--[TTL=3]--------------------------------------------------------------->|
     |<--[Response]-------------------------------------------------------------|
```

### Current Implementation Status

**Note**: This is an early-stage implementation. Currently, the tool performs DNS resolution (translating hostnames to IP addresses). The full traceroute functionality (sending packets with incrementing TTL values) is planned for future development.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.25.3 or later**: [Download Go](https://golang.org/dl/)
  - Verify installation: `go version`
- **Git**: For cloning the repository (if applicable)
- **A Unix-like system** (Linux, macOS, or WSL on Windows) - Required for network operations

## Installation

### Step 1: Clone or Download the Repository

If you have the repository URL:
```bash
git clone <repository-url>
cd traceroute
```

Or if you already have the code:
```bash
cd /path/to/traceroute
```

### Step 2: Install Go Dependencies

This project uses Go's standard library, so no external dependencies need to be installed. However, you should ensure your Go environment is set up correctly:

```bash
# Verify Go is installed
go version

# Check that GOPATH/GOROOT are set (usually automatic)
go env GOPATH
```

### Step 3: Verify Installation

Run a quick test to ensure everything is set up:
```bash
go run ./cmd/traceroute google.com
```

You should see output like:
```
Resolved google.com -> 142.250.191.14
```

## Usage

### Basic Usage

Run traceroute with a hostname or IP address:

```bash
# Using a hostname
go run ./cmd/traceroute google.com

# Using an IP address directly
go run ./cmd/traceroute 8.8.8.8
```

### Using the Makefile

The project includes a Makefile for convenience:

```bash
# Build the binary
make build

# Run the program (pass arguments after make)
make run google.com

# Clean build artifacts
make clean
```

### Using Build Scripts Directly

```bash
# Build the binary
./scripts/build.sh

# Run the program
./scripts/run.sh google.com
```

### Running the Compiled Binary

After building:

```bash
# Build first
make build

# Run the binary directly
./bin/traceroute google.com
```

## Project Structure

```
traceroute/
├── cmd/
│   └── traceroute/
│       └── main.go          # Main entry point - handles CLI arguments and orchestration
├── internal/
│   └── utils/
│       └── resolve.go       # DNS resolution utilities
├── scripts/
│   ├── build.sh             # Build script
│   └── run.sh               # Run script
├── bin/                     # Compiled binaries (created after build)
├── go.mod                   # Go module definition
├── Makefile                 # Build automation
└── README.md                # This file
```

### File Descriptions

- **`cmd/traceroute/main.go`**: The main application entry point. Handles command-line argument parsing and coordinates the traceroute process.
- **`internal/utils/resolve.go`**: Contains DNS resolution logic. Translates human-readable hostnames (like "google.com") into IP addresses.
- **`scripts/build.sh`**: Compiles the Go code into a standalone binary executable.
- **`scripts/run.sh`**: Convenience script to run the program without building first.
- **`Makefile`**: Provides convenient commands (`make build`, `make run`, `make clean`) for common tasks.

## How It Works (Technical Details)

### DNS Resolution (Current Implementation)

The current implementation focuses on **DNS (Domain Name System) resolution**, which is the first step in any network communication:

#### What is DNS?

DNS is like the internet's phone book. When you type "google.com" into your browser, DNS translates that human-readable name into an IP address like "142.250.191.14" that computers use to route packets.

#### The Resolution Process

1. **Input**: User provides a hostname (e.g., "google.com")
2. **DNS Query**: The program uses Go's `net.LookupIP()` function to query DNS servers
3. **Response Processing**: DNS servers return all IP addresses associated with the hostname (both IPv4 and IPv6)
4. **IPv4 Selection**: The code filters for IPv4 addresses (32-bit addresses like 192.168.1.1)
5. **Output**: Returns the first IPv4 address found

#### Code Flow

```
main.go
  │
  ├─> Parse command-line arguments
  │
  ├─> Call utils.Resolve(hostname)
  │     │
  │     └─> resolve.go
  │           │
  │           ├─> net.LookupIP(hostname)
  │           │     └─> Queries DNS servers
  │           │
  │           ├─> Filter for IPv4 addresses
  │           │
  │           └─> Return IP address string
  │
  └─> Display result
```

### Future: Full Traceroute Implementation

When fully implemented, the traceroute process will work as follows:

1. **DNS Resolution**: Convert hostname to IP address (current implementation)
2. **Packet Creation**: Create UDP or ICMP packets with incrementing TTL values
3. **Packet Sending**: Send packets starting with TTL=1
4. **Response Handling**: Receive ICMP "Time Exceeded" messages from intermediate routers
5. **Hop Discovery**: Each response reveals a router in the path
6. **Latency Measurement**: Measure round-trip time for each hop
7. **Path Display**: Show the complete path with timing information

### Network Concepts Explained

#### IP Addresses

- **IPv4**: 32-bit addresses written as four numbers (0-255) separated by dots
  - Example: `192.168.1.1`, `8.8.8.8`, `142.250.191.14`
  - Format: `xxx.xxx.xxx.xxx` where each xxx is 0-255
- **IPv6**: 128-bit addresses written in hexadecimal
  - Example: `2001:0db8:85a3:0000:0000:8a2e:0370:7334`
  - Designed to replace IPv4 due to address exhaustion

#### DNS Records

- **A Record**: Maps a hostname to an IPv4 address
- **AAAA Record**: Maps a hostname to an IPv6 address
- **CNAME Record**: Maps a hostname to another hostname (alias)

#### Routers and Hops

- **Router**: A network device that forwards packets between networks
- **Hop**: Each router a packet passes through is called a "hop"
- **Routing Table**: Routers use routing tables to decide where to send packets next

#### ICMP (Internet Control Message Protocol)

- **Purpose**: Used for network diagnostics and error reporting
- **ICMP Time Exceeded**: Sent by routers when a packet's TTL reaches zero
- **ICMP Echo Request/Reply**: Used by ping to test connectivity

## Building from Source

### Manual Build

```bash
# Navigate to project directory
cd traceroute

# Build the binary
go build -o ./bin/traceroute ./cmd/traceroute

# The binary will be created at ./bin/traceroute
```

### Using the Makefile

```bash
# Build
make build

# The binary will be at ./bin/traceroute
```

### Build Script

```bash
# Make script executable (first time only)
chmod +x scripts/build.sh

# Run build script
./scripts/build.sh
```

## Development

### Running Tests

Currently, the project doesn't include tests, but you can test manually:

```bash
# Test DNS resolution
go run ./cmd/traceroute google.com
go run ./cmd/traceroute github.com
go run ./cmd/traceroute 8.8.8.8
```

### Code Style

This project follows Go's standard formatting conventions:

```bash
# Format code
go fmt ./...

# Check for common errors
go vet ./...
```

### Adding Features

To extend this project, you might want to:

1. **Implement Full Traceroute**: Add packet sending with incrementing TTL values
2. **Add IPv6 Support**: Extend DNS resolution to handle IPv6 addresses
3. **Add Timing Information**: Measure and display latency for each hop
4. **Add Options**: Support command-line flags (e.g., `-n` for numeric output, `-m` for max hops)
5. **Add Error Handling**: Better handling of network errors and timeouts

## Troubleshooting

### Common Issues

**Problem**: `command not found: go`
- **Solution**: Install Go from [golang.org/dl](https://golang.org/dl/)

**Problem**: `DNS lookup failed`
- **Solution**: Check your internet connection and DNS settings
- Try: `ping google.com` to verify connectivity

**Problem**: `permission denied` when running scripts
- **Solution**: Make scripts executable: `chmod +x scripts/*.sh`

**Problem**: `no A records found`
- **Solution**: The hostname might only have IPv6 addresses, or the hostname doesn't exist
- Try: `nslookup <hostname>` to verify DNS records

## Learning Resources

To deepen your understanding of networking:

- **Computer Networks**: Learn about the OSI model, TCP/IP, and routing protocols
- **DNS**: Understand how DNS resolution works and DNS record types
- **ICMP**: Study Internet Control Message Protocol and its uses
- **IP Protocol**: Learn about IP headers, TTL, and packet structure
- **Network Tools**: Explore `ping`, `nslookup`, `dig`, and standard `traceroute`

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here if applicable]

---

**Note**: This is an educational project demonstrating network programming concepts. For production use, consider using established tools like the system `traceroute` command or `mtr` (My Traceroute).


