default:
	make dev

docker:
	@echo "🐋 Containerizing Service"
	docker build -t etl-magic-eden-collection-market . 
	docker run -t etl-magic-eden-collection-market -p 8080:8080

dev:
	go run ./server/main.go