run	:
	go run main.go

db	: 
	docker run --name postgresdb -p 4321:5432 -e POSTGRES_PASSWORD=secret -d postgres:latest

rundb	:
	docker start postgresdb

stopdb	:
	docker stop postgresdb

redis : 
	docker run --name my-redis -p 6379:6379 -d redis

runredis:
	docker start my-redis

runrediscli:
	docker exec -it my-redis redis-cli