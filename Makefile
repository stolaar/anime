target: dev-api dev-client

migrate:
	- cd api && go run migrations/entry.go

install-api:
	- cd api && go mod tidy

install-client:
	- cd client && yarn

dev-api:
	- cd api && go run cmd/main.go

dev-client:
	- cd frontend && yarn dev
