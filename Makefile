main: main.go
	CGO_ENABLED=0 go build main.go

docker: main
	docker build -t ghcr.io/chanwit/spike-bucket .
	docker push ghcr.io/chanwit/spike-bucket
