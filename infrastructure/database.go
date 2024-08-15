package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maooz4426/SDVX-Database/domain"
	"log"
)

func CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS musics (
    	name TEXT NOT NULL,
    	artist TEXT NOT NULL
);`
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("database can't create table", err)
	}
}

func InsertMusic(db *sql.DB, music domain.Music) {
	query := `INSERT INTO musics (name, artist) VALUES (?, ?);`

	log.Println(music.Name)

	_, err := db.Exec(query, music.Name, music.Artist)

	if err != nil {
		log.Fatal("database can't insert music", err)
	}
}
