dev-start:
	docker-compose -f ./docker/docker-compose.yml --env-file ./docker/.env up -d 

dev-stop:
	docker-compose -f ./docker/docker-compose.yml  --env-file ./docker/.env down 

tests:
	go test -v -cover ./...

unit-tests:
	go test --tags=unit ./...

e2e:
	go test --tags=integration ./...

mockgen:
	mockgen -package contact -destination contact/mockstore.go github.com/valverde.thiago/go-agenda-api/contact Store

server:
	go run main.go

.PHONY: dev-start dev-stop tests unit-tests e2e server mockgen