build:
	@go build -o bin/fintechVault

run : build 
	@./bin/fintechVault

test: 
	@go test -v ./...

docker:
	@docker run --name fintech -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=ethereumsolana -d postgres:16-alpine

createdb:
	docker exec -it fintech createdb --username=root --owner=root fintechdb