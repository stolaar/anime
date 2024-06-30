target: dev-backend dev-frontend

dev-backend:
	- cd api && go run cmd/main.go
dev-frontend:
	- cd frontend && yarn dev
