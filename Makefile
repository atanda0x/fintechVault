build:
	@go build -o bin/fintechVault

run : build 
	@./bin/fintechVault

test: 
	@go test -v ./...