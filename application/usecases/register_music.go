package usecases

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/maooz4426/SDVX-Database/domain/model"
	"github.com/maooz4426/SDVX-Database/domain/repository"
	"github.com/sclevine/agouti"
	"log"
	"strconv"
	"strings"
	"time"
)

// 楽曲登録を定義
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

	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--no-sandbox",
			"--disable-dev-shm-usage",
		})

	driver := agouti.ChromeDriver(options)

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
		log.Fatal("Failed to navigate: %v", err)
	}

	//pageの枚数を確認する
	maxPageValue, err := page.All("select#search_page option").Count()

	fmt.Println("maxPageValue", maxPageValue)

	selectBox := page.FindByName("page")
	for i := 0; i < maxPageValue; i++ {
		num := strconv.Itoa(maxPageValue - i)
		err = selectBox.Select(num)
		if err != nil {
			log.Printf("Failed to select: %v", err)

		}

		time.Sleep(3 * time.Second)
		//err = page.SetPageLoad(30)
		if err != nil {
			log.Printf("Failed to set page: %v", err)
		}

		content, err := page.HTML()
		if err != nil {
			log.Printf("Failed to get html: %v", err)
		}

		//fmt.Println("content:", content)

		reader := strings.NewReader(content)
		doc, _ := goquery.NewDocumentFromReader(reader)

		fmt.Println(page, "scan start")

		if err != nil {
			log.Fatal(err)
		}

		var music model.Music
		var level model.Level
		if doc.Find(".music").Length() == 0 {
			fmt.Println(".music要素が見つかりませんでした")
		}

		doc.Find(".music").Each(func(i int, s *goquery.Selection) {
			fmt.Printf("音楽要素 %d が見つかりました\n", i)

			s.Find(".info p").Each(func(i int, info *goquery.Selection) {
				fmt.Printf("info要素 %d が見つかりました: %s\n", i, info.Text())
				if i == 0 {
					music.MusicName = info.Text()
				} else if i == 1 {
					music.Composer = info.Text()
				}
			})

			//fmt.Println("music:", music)

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

			fmt.Println(page, " finished")
		})

		time.Sleep(5 * time.Second)
		////test用
		//if i == 0 {
		//	break
		//}

	}

	fmt.Println("register finished")
	return nil
}
