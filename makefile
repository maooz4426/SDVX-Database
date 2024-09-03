register:
	#go run ./app/cmd/main.go register
	docker compose exec server ./cli_tool register

api:
	#go run ./app/api/main.go
	docker compose up --build


migration/up:
	migrate -path migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/sdvx_db?parseTime=true&charset=utf8mb4" up

migration/down:
	migrate -path migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/sdvx_db?parseTime=true&charset=utf8mb4" down