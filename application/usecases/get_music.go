package usecases

import (
	"context"
	"fmt"
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

func (g *GetMusicUseCase) GetMusicData(ctx context.Context, key string) ([]model.MusicData, error) {

	musicIdList, err := g.musicRepo.SearchMusicData(ctx, key)

	if err != nil {
		return nil, err
	}

	var musicList []model.MusicData

	for _, musicId := range musicIdList {
		musics, err := g.musicRepo.GetMusicData(ctx, musicId)
		fmt.Println(musics)
		for _, music := range musics {
			musicList = append(musicList, music)
		}
		if err != nil {
			return nil, err
		}
	}

	return musicList, nil

}
