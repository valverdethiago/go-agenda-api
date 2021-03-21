dev-start:
	docker-compose -f ./docker/docker-compose.yml  up -d 

dev-stop:
	docker-compose -f ./docker/docker-compose.yml  down 

stack-start:
	docker-compose -f ./docker/docker-compose-full.yml  up -d --force-recreate --build backend

stack-stop:
	docker-compose -f ./docker/docker-compose-full.yml  down 

unit:
	cd src;\
	go test -v -tags=unit -cover ./... 

e2e:
	cd src;\
	go test -v -tags=integration -cover ./... 

tests:
	cd src;\
	go test -v -tags=integration,unit -cover ./...

mockgen:
	cd src;\
	mockgen -package contact -destination contact/mockstore.go github.com/valverde.thiago/go-agenda-api/contact Store

server:
	cd src;\
	go run main.go

build:
	go build -o ./docker/bin/agenda

.PHONY: dev-start dev-stop stack-start stack-stop tests unit e2e server mockgen