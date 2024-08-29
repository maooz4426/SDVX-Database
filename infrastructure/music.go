package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maooz4426/SDVX-Database/domain/model"
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

func NewMusicRepository(db *sql.DB) *MusicRepository {
	return &MusicRepository{db}
}

//mはメソッドレシーバーのこと
//ctxでキャンセル情報を共有する

// 楽曲登録
func (m *MusicRepository) RegisterMusic(ctx context.Context, music model.Music) error {

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

	_, err = tx.ExecContext(ctx, query, music.MusicName, music.Composer, time.Now(), time.Now(), music.MusicName, music.Composer)

	if err != nil {
		tx.Rollback()
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
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	query := `INSERT INTO levels (music_id, level_name, level_value, created_at, updated_at)
SELECT ?, ?, ?, ?, ? WHERE NOT EXISTS (SELECT 1 FROM levels WHERE music_id = ? AND level_name = ? AND level_value = ?);`

	_, err = tx.ExecContext(ctx, query, musicID, level.LevelName, level.LevelValue, time.Now(), time.Now(), musicID, level.LevelName, level.LevelValue)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to add level: %s", err)
	}

	return nil
}

func (m *MusicRepository) GetMusicData(ctx context.Context, musicID string) (musics []model.MusicData, err error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	fmt.Println("id:", musicID)

	query := `SELECT m.music_name, m.composer,l.level_name,l.level_value FROM musics as m LEFT OUTER JOIN levels as l ON m.music_id = l.music_id WHERE m.music_id = ?;`

	rows, err := tx.QueryContext(ctx, query, musicID)
	//fmt.Println(rows)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to get music data: %s", err)
	}

	for rows.Next() {
		var musicName, composer, levelName string
		var levelValue int
		err := rows.Scan(&musicName, &composer, &levelName, &levelValue)
		fmt.Println(musicName, composer, levelName, levelValue)
		if err != nil {
			tx.Rollback()
		}
		musicData := model.MusicData{
			MusicName:  musicName,
			Composer:   composer,
			LevelName:  levelName,
			LevelValue: levelValue,
		}

		fmt.Println(musicData)
		musics = append(musics, musicData)
	}

	return musics, nil

}

func (m *MusicRepository) SearchMusicData(ctx context.Context, key string) ([]string, error) {
	keyword := "%" + key + "%"
	//fmt.Println(keyword)

	tx, err := m.db.Begin()

	query := `SELECT music_id from musics WHERE music_name LIKE ? OR composer LIKE ?;`

	//query := `SELECT m.music_id FROM musics as m LEFT OUTER JOIN levels as l ON m.music_id = l.music_id WHERE music_name LIKE ? OR music_name LIKE ?;`

	rows, err := tx.QueryContext(ctx, query, keyword, keyword)

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to search music data: %s", err)
	}

	var musicIDList []string

	for rows.Next() {
		var musicID string
		err := rows.Scan(&musicID)

		//fmt.Println(musicID)

		musicIDList = append(musicIDList, musicID)

		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to search music data: %s", err)
		}
	}

	return musicIDList, nil

}
