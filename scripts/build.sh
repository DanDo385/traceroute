#!/usr/bin/env bash
set -euo pipefail
echo "Building traceroute CLI..."
go build -o ./bin/traceroute ./cmd/traceroute
echo "Build complete! Binary at ./bin/traceroute"
