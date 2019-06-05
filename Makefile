all: deployer

deployer:
	@echo "build deployer"
	go build -o bin/deployer cmd/main.go

