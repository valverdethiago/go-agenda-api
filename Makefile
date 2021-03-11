dev-start:
	docker-compose -f ./docker/docker-compose.yml --env-file ./docker/.env up -d 

dev-stop:
	docker-compose -f ./docker/docker-compose.yml  --env-file ./docker/.env down 

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: dev-start dev-stop test server