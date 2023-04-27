# Example Go project to use SCOW API and Hook

```bash
# Generate proto files
make protos

# Download modules
go mod tidy

# Run Client call that calls SCOW API
# Make sure SCOW mis-server is listening on 192.168.88.100:7571
# Run
make run-api

# Run Hook server at 5000
make run-hook
```