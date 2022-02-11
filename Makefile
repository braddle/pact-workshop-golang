up:
	docker-compose up -d

down:
	docker-compose down

jump-to-consumer:
	docker-compose exec consumer bash

jump-to-producer:
	docker-compose exec producer bash
