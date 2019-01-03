run:
	docker run --rm -p 8700:8700 -v "${PWD}":/usr/src/chat monorepo/chat go run ./server

build:
	docker build -t "monorepo/chat:1.0.0" -t monorepo/chat:latest .