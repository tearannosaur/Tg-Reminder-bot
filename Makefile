service-up:
	docker compose up -d

migrate-up:
	docker compose run --rm migrate -path /migrations \
	-database "postgres://user:dsakl@db:5432/reminder?sslmode=disable" up

migrate-down:
	docker compose run --rm migrate -path /migrations \
	-database "postgres://user:dsakl@db:5432/reminder?sslmode=disable" down 1

rebuild:
	docker compose down && \
	docker compose build 