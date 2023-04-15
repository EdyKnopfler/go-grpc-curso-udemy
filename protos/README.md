
# Protocol Buffers

Arquivos `*.pb.go` gerados a partir de `signature.proto` usando:

```bash
PATH="$PATH:$(go env GOPATH)/bin" protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    signature.proto
```