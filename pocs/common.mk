.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: clean
clean:
	docker compose down -v

.PHONY: build
build: 
	docker compose up --build

.PHONY: watch
watch:
	docker compose watch

.PHONY: health
health:
	docker compose ps

.PHONY: logs
logs:
	docker compose logs --follow --tail 1000
