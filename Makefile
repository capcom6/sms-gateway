project_name = sms-gateway
image_name = capcom6/$(project_name):latest

extension=
ifeq ($(OS),Windows_NT)
	extension = .exe
endif

init:
	go mod download

init-dev: init
	go install github.com/cosmtrek/air@latest \
		&& go install github.com/swaggo/swag/cmd/swag@latest \
		&& go install github.com/pressly/goose/v3/cmd/goose@latest

air:
	air

db-upgrade:
	go run ./cmd/$(project_name)/main.go db:migrate

db-upgrade-raw:
	go run ./cmd/$(project_name)/main.go db:auto-migrate
	
run:
	go run cmd/$(project_name)/main.go

test:
	go test -cover ./...

build:
	go build -o tmp/$(project_name) ./cmd/$(project_name)
	
install:
	go install ./cmd/$(project_name)

docker-build:
	docker build -f build/package/Dockerfile -t $(image_name) --build-arg APP=$(project_name) .

docker:
	docker-compose -f deployments/docker-compose/docker-compose.yml up --build

docker-dev:
	docker-compose -f deployments/docker-compose/docker-compose.dev.yml up --build

api-docs:
	swag fmt -g ./cmd/$(project_name)/main.go \
		&& swag init -g ./cmd/$(project_name)/main.go -o ./api

view-docs:
	php -S 127.0.0.1:8080 -t ./api

clean:
	docker-compose -f deployments/docker-compose/docker-compose.yml down --volumes

.PHONY: init init-dev air db-upgrade db-upgrade-raw run test build install docker docker-dev api-docs view-docs clean