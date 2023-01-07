.PHONY: init test build serve clean pack run

include .env
export $(shell sed 's/=.*//' .env)

ifeq ($(HOST_PORT),)
    HOST_PORT := 8080
endif

init:
	@go mod download

test: init
	@go test ./... -count=1 -cover

build: init
	@rm -rf ./dist
	@mkdir dist
	@cp -r api ./dist/api
	@go build -o ./dist/flyer .	

serve: build
	@cd dist && ./flyer 

clean:
	@rm ./dist/ -rf
	@-docker container rm flyer
	@-docker compose down
	@-docker image rm flyer

pack:
	@docker build --build-arg APP_NAME=flyer -t flyer .

run:
	@-docker container rm flyer
	@docker run --name flyer -p $(HOST_PORT):$(HOST_PORT) flyer:latest

compose:
	@docker compose up -d
	@sleep 5s
	@open http://localhost:7080