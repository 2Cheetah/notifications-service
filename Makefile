.PHONY: run
.DEFAULT: run

run:
	go run main.go

compose-build:
	docker compose -f docker-compose.yml build

compose-up:
	docker compose -f docker-compose.yml up

compose-down:
	docker compose -f docker-compose.yml down
