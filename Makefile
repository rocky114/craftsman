sqlc:
	sqlc generate

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o craftman cmd/server/main.go
