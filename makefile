BUILDHASH = $(shell git rev-parse --verify HEAD | cut -c 1-7)
VERSION = 1.0.0

# go build
gobuild_http:
	@go build -v -o simple-service-http cmd/http/*.go

# go run
gorun_http:
	make gobuild_http
	@./simple-service-http

docker-up:
	@docker-compose up -d

docker-down:
	@ rm -r -f ./postgres-data
	@docker-compose down
	@docker volume rm "simple-service_db"