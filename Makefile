dev-start:
	docker-compose -f ./docker/docker-compose.yml  up -d 

dev-stop:
	docker-compose -f ./docker/docker-compose.yml  down 

tests:
	cd src;\
	go test -v -cover ./...

mockgen:
	cd src;\
	mockgen -package contact -destination contact/mockstore.go github.com/valverde.thiago/go-agenda-api/contact Store

server:
	cd src;\
	go run main.go

build:
	go build -o ./docker/bin/agenda

.PHONY: dev-start dev-stop tests unit-tests e2e server mockgen