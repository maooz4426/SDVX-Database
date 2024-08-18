package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maooz4426/SDVX-Database/domain"
	"time"
)

type MusicRepository struct {
	db *sql.DB
}

func NewMusicRepository(db *sql.DB) *MusicRepository {
	return &MusicRepository{db}
}

//func CreateTable(db *sql.DB) {
//
//	//存在したら削除
//	dropquery := `DROP TABLE IF EXISTS musics;`
//	_, err := db.Exec(dropquery)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	query := `CREATE TABLE IF NOT EXISTS musics (
//    	music_id INTERGER NOT NULL PRIMARY KEY,
//    	music_name TEXT NOT NULL,
//    	composer TEXT NOT NULL,
//    	createdAt TIMESTAMP NOT NULL,
//    	updatedAt TIMESTAMP NOT NULL
//);`
//	_, err = db.Exec(query)
//
//	if err != nil {
//		log.Fatal("database can't create table", err)
//	}
//}

//mはメソッドレシーバーのこと

//ctxでキャンセル情報を共有する

// 楽曲登録
func (m *MusicRepository) RegisterMusic(ctx context.Context, music domain.Music) error {
	query := `INSERT INTO musics (music_name, composer,createdAt,updatedAt) VALUES (?, ?,?,?);`

	//log.Println(music.MusicName)

	_, err := m.db.ExecContext(ctx, query, music.MusicName, music.Composer, time.Now(), time.Now())

	if err != nil {
		//log.Fatal("database can't insert music", err)
		return fmt.Errorf("Error register music :%s", err)
	}

	return nil
}

func (m *MusicRepository) GetMusicID(ctx context.Context, music domain.Music) (int, error) {
	var musicID int
	query := `SELECT music_id FROM musics WHERE music_name = ? AND composer = ?;`

	err := m.db.QueryRowContext(ctx, query, music.MusicName, music.Composer).Scan(&musicID)
	if err != nil {

		return 0, fmt.Errorf("failed to get music ID: %s", err)
	}

	return musicID, nil
}

func (m *MusicRepository) RegisterLevel(ctx context.Context, musicID int, level domain.Level) error {
	query := `INSERT INTO levels(music_id,level_name, level_value,created_at,updated_at) VALUES (?, ?,?,?,?);`

	_, err := m.db.ExecContext(ctx, query, musicID, level.LevelName, level.LevelValue, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to add level: %s", err)
	}

	return nil
}
