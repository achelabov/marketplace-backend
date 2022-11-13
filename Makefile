.PHONY:
.SILENT:

test:
	go test -v -cover -race ./...

build:
	go build -o .bin/app ./cmd/app/main.go

run: build
	docker-compose up --remove-orphans

stop:
	docker-compose down --remove-orphans

run_local: build
	./.bin/app