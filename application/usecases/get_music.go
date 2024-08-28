package usecases

import (
	"context"
	"github.com/maooz4426/SDVX-Database/domain/model"
	"github.com/maooz4426/SDVX-Database/domain/repository"
)

type GetMusicUseCaser interface {
	//可変長引数を使う
	// ...が可変長
	//空のインターフェースを作成するとなんでも入れることができる
	GetMusicData(ctx context.Context, musicID string) (model.MusicData, error)
	SearchMusicData(ctx context.Context, args ...interface{}) ([]model.MusicData, error)
}

type GetMusicUseCase struct {
	musicRepo repository.MusicRepositoryImpl
}

func NewGetMusicUseCase(m repository.MusicRepositoryImpl) *GetMusicUseCase {
	return &GetMusicUseCase{
		musicRepo: m,
	}
}

func (g *GetMusicUseCase) GetMusicData(ctx context.Context, musicID string) ([]model.MusicData, error) {
	musics, err := g.musicRepo.GetMusicData(ctx, musicID)
	if err != nil {
		return nil, err
	}

	//jsonData, err := json.Marshal(musics)
	if err != nil {
		return nil, err
	}

	return musics, nil

	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(jsonData)
}
