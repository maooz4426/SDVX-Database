register:
	go run ./app/cmd/main.go register

migration up:
	migrate -path migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/sdvx_db?parseTime=true&charset=utf8mb4" up

migration down:
	migrate -path migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/sdvx_db?parseTime=true&charset=utf8mb4" down