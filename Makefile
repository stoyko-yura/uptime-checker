run-all:
	docker-compose up -d
	cmd /c start cmd /k "cd server && go run cmd/api/main.go"
	cmd /c start cmd /k "cd client && npm run dev"
stop-all:
	docker-compose down
	@taskkill /F /IM main.exe /T 2>nul || echo
	@taskkill /F /IM node.exe /T 2>nul || echo