all: trigger deployer

deployer:
	@echo "build deployer"
	go build -o bin/deployer cmd/deployer/main.go

trigger:
	@echo "build trigger"
	go build -o bin/trigger cmd/trigger/main.go

