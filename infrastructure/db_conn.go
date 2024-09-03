package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func NewDBConn() (*sql.DB, error) {
	sqluser := "user"

	sqlpass := "password"

	sqldb := "sdvx_db"

	//dbのセットアップ

	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?parseTime=true&charset=utf8mb4", sqluser, sqlpass, sqldb)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 6; i++ {
		if errcon := db.Ping(); errcon != nil {
			log.Println("Now Trying count:", i)
			log.Println("Failed to connect to database", errcon)
			db, _ = sql.Open("mysql", dsn)
		} else {
			//break
			fmt.Println("Connected to database")
			return db, nil
		}
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Failed to connect to database")

	return nil, err
}
