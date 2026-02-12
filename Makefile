.PHONY: frontend backend prod dev 

frontend:
	cd frontend && npm run build

backend:
	cd backend && go run main.go

prod: frontend
	cd backend && go build -o server && ./server

dev:
	@trap 'kill 0' EXIT; \
	cd frontend && npm run dev & \
	cd backend && go run main.go & \
	wait

