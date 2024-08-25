package usecases

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/maooz4426/SDVX-Database/domain/model"
	"github.com/maooz4426/SDVX-Database/domain/repository"
	"github.com/sclevine/agouti"
	"log"
	"strconv"
	"strings"
	"time"
)

// メソッド抽象化して依存させる
//type MusicRepository interface {
//	RegisterMusic(ctx context.Context, music domain.Music) error
//	GetMusicID(ctx context.Context, music domain.Music) (int, error)
//	RegisterLevel(ctx context.Context, musicID int, level domain.Level) error
//}

type MusicRegisterImpl interface {
	Register(ctx context.Context) error
}

type MusicRegister struct {
	musicRepo repository.MusicRepositoryImpl
}

// インスタンスメソッド
func NewRegisterMusicData(m repository.MusicRepositoryImpl) MusicRegisterImpl {
	return &MusicRegister{
		musicRepo: m,
	}
}

func (a *MusicRegister) Register(ctx context.Context) error {
	var err error

	url := "https://p.eagate.573.jp/game/sdvx/vi/music/index.html"

	driver := agouti.ChromeDriver()

	err = driver.Start()
	if err != nil {
		log.Fatal("driver start error", err)
	}

	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}

	err = page.Navigate(url)
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	//pageの枚数を確認する
	maxPageValue, err := page.All("select#search_page option").Count()

	selectBox := page.FindByName("page")
	for i := 0; i < maxPageValue; i++ {
		num := strconv.Itoa(maxPageValue - i)
		err = selectBox.Select(num)
		if err != nil {
			log.Printf("Failed to select: %v", err)

		}

		content, err := page.HTML()
		if err != nil {
			log.Printf("Failed to get html: %v", err)
		}

		reader := strings.NewReader(content)
		doc, _ := goquery.NewDocumentFromReader(reader)

		if err != nil {
			log.Fatal(err)
		}

		var music model.Music
		var level model.Level
		doc.Find(".music").Each(func(i int, s *goquery.Selection) {
			s.Find(".info p").Each(func(i int, info *goquery.Selection) {
				if i == 0 {
					music.MusicName = info.Text()
					//fmt.Println(music.MusicName)
				} else if i == 1 {
					music.Composer = info.Text()
					//fmt.Println(music.Composer)
				}
			})

			//fmt.Println(music)

			if err := a.musicRepo.RegisterMusic(ctx, music); err != nil {
				log.Printf("Failed to register music %s: %v", music.MusicName, err)
				return
			}

			s.Find(".level p").Each(func(i int, levelV *goquery.Selection) {
				//levelの値取得
				var levelNum int
				levelNum, err = strconv.Atoi(levelV.Text())
				if err != nil {
					log.Fatal(err)
				}
				level.LevelValue = levelNum

				//level名取得
				level.LevelName = levelV.AttrOr("class", "")

				var musicID int
				musicID, err = a.musicRepo.GetMusicID(ctx, music)

				err = a.musicRepo.RegisterLevel(ctx, musicID, level)

				if err != nil {
					log.Fatal(err)
				}

			})
		})

		time.Sleep(2 * time.Second)

		//test用
		if i == 0 {
			break
		}

	}

	return nil
}
