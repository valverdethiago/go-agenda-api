dev-start:
	docker-compose -f ./docker/docker-compose.yml --env-file ./docker/.env up -d 

dev-stop:
	docker-compose -f ./docker/docker-compose.yml  --env-file ./docker/.env down 

test:
	go test -v -cover ./...

mockgen:
	mockgen -package mockdb -destination contact/mock/mockstore.go github.com/valverde.thiago/go-agenda-api/contact Store

server:
	go run main.go

.PHONY: dev-start dev-stop test server mockgen