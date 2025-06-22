.PHONY: run
.DEFAULT: run

run: test
	go run main.go

test:
	go test -v .

compose-build:
	docker compose -f docker-compose.yml build

compose-up:
	docker compose -f docker-compose.yml up

compose-down:
	docker compose -f docker-compose.yml down
