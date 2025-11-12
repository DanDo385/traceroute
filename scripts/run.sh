#!/usr/bin/env bash
set -euo pipefail
echo "Running traceroute CLI..."
go run ./cmd/traceroute "$@"
