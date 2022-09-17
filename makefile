default:
	@echo "craftsman"

## http: 启动http服务
run-http:
	go run main.go

## sql生成
sqlc:
	sqlc generate -f config/sqlc.yaml

help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'