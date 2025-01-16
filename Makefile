run:
	@echo "service is running..."
	@go run main.go

run-three-services:
	@echo "three service is running..."
	@go run main.go 3000&
	@go run main.go 3001&
	@go run main.go 3002

run-three-leaky-services:
	@echo "n service is running..."
	@go run main.go 3000 0&
	@go run main.go 3002 3&
	@go run main.go 3004 4
