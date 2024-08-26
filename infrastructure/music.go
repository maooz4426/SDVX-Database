package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maooz4426/SDVX-Database/domain/model"
	"github.com/maooz4426/SDVX-Database/domain/repository"
	"time"
)

type MusicRepository struct {
	db *sql.DB
}

type MusicRepositoryImpl interface {
	RegisterMusic(ctx context.Context, music model.Music) error
	GetMusicID(ctx context.Context, music model.Music) (int, error)
	RegisterLevel(ctx context.Context, musicID int, level model.Level) error
}

func NewMusicRepository(db *sql.DB) repository.MusicRepositoryImpl {
	return &MusicRepository{db}
}

//mはメソッドレシーバーのこと
//ctxでキャンセル情報を共有する

// 楽曲登録
func (m *MusicRepository) RegisterMusic(ctx context.Context, music model.Music) error {
	//query := `INSERT INTO musics (music_name, composer,createdAt,updatedAt) VALUES (?, ?,?,?);`

	//	query := `INSERT INTO musics (music_name, composer,createdAt,updatedAt)
	//SELECT ?,?,?,? WHERE NOT EXISTS (SELECT 1 FROM musics WHERE music_name = ? AND composer = ?)`

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO musics (music_name, composer, createdAt, updatedAt) 
SELECT ?, ?, ?, ?
WHERE NOT EXISTS (
    SELECT 1 
    FROM musics 
    WHERE  (music_name = ? AND composer = ?)
);`

	_, err = m.db.ExecContext(ctx, query, music.MusicName, music.Composer, time.Now(), time.Now(), music.MusicName, music.Composer)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Error register music :%s", err)
	}

	//if err != nil {
	//	//log.Fatal("database can't insert music", err)
	//
	//}

	return nil
}

func (m *MusicRepository) GetMusicID(ctx context.Context, music model.Music) (int, error) {
	var musicID int
	query := `SELECT music_id FROM musics WHERE music_name = ? AND composer = ?;`

	err := m.db.QueryRowContext(ctx, query, music.MusicName, music.Composer).Scan(&musicID)
	if err != nil {

		return 0, fmt.Errorf("failed to get music ID: %s", err)
	}

	return musicID, nil
}

func (m *MusicRepository) RegisterLevel(ctx context.Context, musicID int, level model.Level) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	query := `INSERT INTO levels (music_id, level_name, level_value, created_at, updated_at)
SELECT ?, ?, ?, ?, ? WHERE NOT EXISTS (SELECT 1 FROM levels WHERE music_id = ? AND level_name = ? AND level_value = ?);`

	_, err = m.db.ExecContext(ctx, query, musicID, level.LevelName, level.LevelValue, time.Now(), time.Now(), musicID, level.LevelName, level.LevelValue)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to add level: %s", err)
	}

	return nil
}
