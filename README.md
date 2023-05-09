# Example Go project to use SCOW API and Hook

## Prerequisites

- [Go](https://go.dev/) 1.20
- [Buf CLI](https://buf.build/product/cli/)

## Commands

```bash

# Use a specific SCOW API version to generate proto files
# starting from v0.3.0
SCOW_API_VERSION=v0.3.0 make protos

# Download modules
go mod tidy

# Run Client call that calls SCOW API
# Make sure SCOW mis-server is listening on 192.168.88.100:7571
# Run
make run-api

# If your API has setup token authentication, set your token by env SCOW_API_TOKEN
SCOW_API_TOKEN={Token} make run api

# Run Hook server at 5000
make run-hook
```