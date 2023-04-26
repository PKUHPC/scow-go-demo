# Example Go project to use SCOW API

```bash
# Generate proto files
make protos

# Download modules
go mod tidy

# Make sure SCOW mis-server is listening on 192.168.88.100:7571

# Run
go run main.go
```