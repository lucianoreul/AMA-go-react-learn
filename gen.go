package gen

//go:generate go run ./cmd/tools/terndotenv/main.go tun go generate ./... | run go generate
//go:generate sqlc generate -f ./internal/store/pgstore/sqlc.yaml
