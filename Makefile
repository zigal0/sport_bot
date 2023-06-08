# App
.PHONY: run
run:
	go run cmd/sport_bot/main.go

## Docker
.PHONY: compose-up
compose-up:
	docker compose -p sport_bot -f ./build/docker-compose.yml up -d

.PHONY: compose-down
compose-down:
	docker compose -p sport_bot -f ./build/docker-compose.yml down

.PHONY: compose-rm
compose-rm:
	docker compose -p sport_bot -f ./build/docker-compose.yml rm -fvs  #f - force, v - any anonymous volumes, s - stop

.PHONY: compose-rs
compose-rs:
	make compose-rm
	make compose-up

.PHONY: run-full
	run-full:
	make compose-up
	echo "Waiting for DBs to start"
	sleep 10
	make run