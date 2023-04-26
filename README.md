# Example Go project to use SCOW API

```bash
# Generate proto files
make protos

# Make sure SCOW mis-server is listening on 192.168.88.100:7571

# Download modules
go mod tidy

# Run
go run main.go
```