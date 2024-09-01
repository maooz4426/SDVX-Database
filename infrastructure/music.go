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
	query := `INSERT INTO musics (music_name, composer, createdAt, updatedAt) 
               SELECT ?, ?, ?, ?
               WHERE NOT EXISTS (
                   SELECT 1 
                   FROM musics 
                   WHERE (music_name = ? AND composer = ?)
               );`

	_, err := m.db.ExecContext(ctx, query, music.MusicName, music.Composer, time.Now(), time.Now(), music.MusicName, music.Composer)
	if err != nil {
		return fmt.Errorf("Error register music: %w", err)
	}

	fmt.Println("register:", music.MusicName)
	return nil
}

func (m *MusicRepository) GetMusicID(ctx context.Context, music model.Music) (int, error) {
	var musicID int
	query := `SELECT music_id FROM musics WHERE music_name = ? AND composer = ?;`

	err := m.db.QueryRowContext(ctx, query, music.MusicName, music.Composer).Scan(&musicID)
	if err != nil {
		return 0, fmt.Errorf("failed to get music ID: %w", err)
	}

	return musicID, nil
}

func (m *MusicRepository) RegisterLevel(ctx context.Context, musicID int, level model.Level) error {
	query := `INSERT INTO levels (music_id, level_name, level_value, created_at, updated_at)
               SELECT ?, ?, ?, ?, ? 
               WHERE NOT EXISTS (
                   SELECT 1 
                   FROM levels 
                   WHERE music_id = ? AND level_name = ? AND level_value = ?
               );`

	_, err := m.db.ExecContext(ctx, query, musicID, level.LevelName, level.LevelValue, time.Now(), time.Now(), musicID, level.LevelName, level.LevelValue)
	if err != nil {
		return fmt.Errorf("failed to add level: %w", err)
	}

	return nil
}

func (m *MusicRepository) GetMusicData(ctx context.Context, musicID string) ([]model.MusicData, error) {
	query := `SELECT m.music_name, m.composer, l.level_name, l.level_value 
              FROM musics as m 
              LEFT OUTER JOIN levels as l ON m.music_id = l.music_id 
              WHERE m.music_id = ?;`

	rows, err := m.db.QueryContext(ctx, query, musicID)
	if err != nil {
		return nil, fmt.Errorf("failed to get music data: %w", err)
	}
	defer rows.Close()

	var musics []model.MusicData
	for rows.Next() {
		var musicName, composer, levelName string
		var levelValue int
		if err := rows.Scan(&musicName, &composer, &levelName, &levelValue); err != nil {
			return nil, fmt.Errorf("failed to scan music data: %w", err)
		}
		musicData := model.MusicData{
			MusicName:  musicName,
			Composer:   composer,
			LevelName:  levelName,
			LevelValue: levelValue,
		}
		musics = append(musics, musicData)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating music data rows: %w", err)
	}

	return musics, nil
}

func (m *MusicRepository) SearchMusicData(ctx context.Context, key string) ([]string, error) {
	keyword := "%" + key + "%"
	query := `SELECT music_id FROM musics WHERE music_name LIKE ? OR composer LIKE ?;`

	rows, err := m.db.QueryContext(ctx, query, keyword, keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to search music data: %w", err)
	}
	defer rows.Close()

	var musicIDList []string
	for rows.Next() {
		var musicID string
		if err := rows.Scan(&musicID); err != nil {
			return nil, fmt.Errorf("failed to scan music ID: %w", err)
		}
		musicIDList = append(musicIDList, musicID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating search results: %w", err)
	}

	return musicIDList, nil
}
