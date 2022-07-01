SHA=$(shell git rev-parse --short HEAD)

main: main.go
	CGO_ENABLED=0 go build main.go

docker: main
	docker build --no-cache -t ghcr.io/chanwit/spike-bucket:${SHA} .
	docker push ghcr.io/chanwit/spike-bucket:${SHA}
