default:
	@echo "craftsman"

## http: 启动http服务
http:
	go run main.go

help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'