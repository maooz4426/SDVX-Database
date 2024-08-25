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

//	func NewMusicRepository(db *sql.DB) *MusicRepository {
//		return &MusicRepository{db}
//	}
func NewMusicRepository(db *sql.DB) repository.MusicRepositoryImpl {
	return &MusicRepository{db}
}

//mはメソッドレシーバーのこと

//ctxでキャンセル情報を共有する

// 楽曲登録
func (m *MusicRepository) RegisterMusic(ctx context.Context, music model.Music) error {
	query := `INSERT INTO musics (music_name, composer,createdAt,updatedAt) VALUES (?, ?,?,?);`

	//log.Println(music.MusicName)
	//log.Println(music)

	_, err := m.db.ExecContext(ctx, query, music.MusicName, music.Composer, time.Now(), time.Now())

	if err != nil {
		//log.Fatal("database can't insert music", err)
		return fmt.Errorf("Error register music :%s", err)
	}

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
	query := `INSERT INTO levels(music_id,level_name, level_value,created_at,updated_at) VALUES (?, ?,?,?,?);`

	_, err := m.db.ExecContext(ctx, query, musicID, level.LevelName, level.LevelValue, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to add level: %s", err)
	}

	return nil
}
