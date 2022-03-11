start:
	docker compose up -d --force-recreate --build

stop:
	docker-compose down

jump-to-consumer:
	docker-compose exec consumer bash

jump-to-producer:
	docker-compose exec producer bash
