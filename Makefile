all: stop start

stop:
	docker compose down

start:
	docker compose build --no-cache
	docker compose up & 

clean:
	docker ps -a -q | xargs docker rm
	docker images -q | xargs docker rmi -f