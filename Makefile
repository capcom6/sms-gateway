project_name = smsgateway
image_name = cr.selcloud.ru/soft-c/$(project_name):latest

extension=
ifeq ($(OS),Windows_NT)
	extension = .exe
endif

init:
	go mod download \
		&& go install github.com/pressly/goose/v3/cmd/goose@latest

init-dev: init
	go install github.com/cosmtrek/air@latest \
		&& go install github.com/swaggo/swag/cmd/swag@latest

air:
	air

db-upgrade:
	goose up

db-upgrade-raw:
	go run ./cmd/$(project_name)/main.go db-upgrade

test:
	go test ./...

api-docs:
	swag fmt -g ./cmd/$(project_name)/main.go \
		&& swag init -g ./cmd/$(project_name)/main.go -o ./api

view-docs:
	php -S 127.0.0.1:8080 -t ./api

.PHONY: init init-dev air db-upgrade db-upgrade-raw test api-docs view-docs
