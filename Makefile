.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t telegbotyoutube:v0.1 .

start-container:
	docker run --name talegram-bot -p 8183:8183 --env-file .env telegbotyoutube:v0.1

push-dockerhub:
	docker build -t telegbotyoutube:v0.1 .
	docker tag telegbotyoutube:v0.1 mukolla/tgbot:latest
	docker push mukolla/tgbot:latest