target: dev-backend dev-frontend

dev-backend:
	- go run main.go
dev-frontend:
	- cd frontend && yarn dev
