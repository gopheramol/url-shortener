run:
	go run cmd/app/main.go
test:
	go test ./... --cover
build:
	go build cmd/app/main.go
download:
		go mod tidy