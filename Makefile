BINARY_NAME=cippgoapi
build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} cmd/cippgoapi/main.go
test:
	go test ./... -v