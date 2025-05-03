dev:
	go build cmd/server/main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o craftman cmd/server/main.go

sqlc:
	sqlc generate

# 清理构建产物
.PHONY: clean
clean:
	rm -rf main