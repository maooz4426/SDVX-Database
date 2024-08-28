package repository

import (
	"context"
	"github.com/maooz4426/SDVX-Database/domain/model"
)

// ビジネスルールをここで定義
type MusicRepositoryImpl interface {
	RegisterMusic(ctx context.Context, music model.Music) error
	GetMusicID(ctx context.Context, music model.Music) (int, error)
	RegisterLevel(ctx context.Context, musicID int, level model.Level) error
	GetMusicData(ctx context.Context, musicID string) (musics []model.MusicData, err error)
}
