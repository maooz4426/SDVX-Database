package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maooz4426/SDVX-Database/repository"
	"github.com/maooz4426/SDVX-Database/services"
	"log"
	"time"
)

func main() {

	sqluser := "user"

	sqlpass := "password"

	sqldb := "sdvx_db"

	//dbのセットアップ

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true&charset=utf8mb4", sqluser, sqlpass, sqldb)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 6; i++ {
		if errcon := db.Ping(); errcon != nil {
			log.Println("Now Trying count:", i)
			log.Fatal("Failed to connect to database", errcon)
			db, _ = sql.Open("mysql", dsn)
		} else {
			break
		}
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Connected to database")

	rep := repository.NewMusicRepository(db)

	svc := services.NewService(rep)

	ctx := context.Background()

	svc.RegisterMusicData(ctx)

	//テーブル作成
	//infrastructure.CreateTable(db)

	//url := "https://p.eagate.573.jp/game/sdvx/vi/music/index.html"
	//
	//driver := agouti.ChromeDriver()
	//
	//err = driver.Start()
	//if err != nil {
	//	log.Fatal("driver start error", err)
	//}
	//
	//defer driver.Stop()
	//
	//page, err := driver.NewPage(agouti.Browser("chrome"))
	//if err != nil {
	//	log.Printf("Failed to open page: %v", err)
	//}
	//
	//err = page.Navigate(url)
	//if err != nil {
	//	log.Printf("Failed to navigate: %v", err)
	//}
	//
	////pageの枚数を確認する
	//maxPageValue, err := page.All("select#search_page option").Count()
	//
	//selectBox := page.FindByName("page")
	//for i := 0; i < maxPageValue; i++ {
	//	num := strconv.Itoa(maxPageValue - i)
	//	err = selectBox.Select(num)
	//	if err != nil {
	//		log.Printf("Failed to select: %v", err)
	//
	//	}
	//
	//	content, err := page.HTML()
	//	if err != nil {
	//		log.Printf("Failed to get html: %v", err)
	//	}
	//
	//	reader := strings.NewReader(content)
	//	doc, _ := goquery.NewDocumentFromReader(reader)
	//
	//	ctx := context.Background()
	//
	//	//tx, errtx := db.BeginTx(ctx, nil)
	//	//if errtx != nil {
	//	//	log.Printf("Failed to begin tx: %v", errtx)
	//	//}
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	rep := repository.NewMusicRepository(db)
	//
	//	var musics []domain.Music
	//
	//	var music domain.Music
	//	var level domain.Level
	//	doc.Find(".music").Each(func(i int, s *goquery.Selection) {
	//		s.Find(".info p").Each(func(i int, info *goquery.Selection) {
	//			if i == 0 {
	//				music.MusicName = info.Text()
	//				//fmt.Println(music.MusicName)
	//			} else if i == 1 {
	//				music.Composer = info.Text()
	//				//fmt.Println(music.Composer)
	//			}
	//		})
	//
	//		//fmt.Println(music)
	//
	//		err := rep.RegisterMusic(ctx, music)
	//		if err != nil {
	//			log.Fatal("Failed to register music:", err)
	//		}
	//
	//		s.Find(".level p").Each(func(i int, levelV *goquery.Selection) {
	//			//levelの値取得
	//			var levelNum int
	//			levelNum, err = strconv.Atoi(levelV.Text())
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//			level.LevelValue = levelNum
	//
	//			//level名取得
	//			level.LevelName = levelV.AttrOr("class", "")
	//
	//			var musicID int
	//			musicID, err = rep.GetMusicID(ctx, music)
	//
	//			//fmt.Println(musicID)
	//
	//			err = rep.RegisterLevel(ctx, musicID, level)
	//
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//
	//		})
	//	})
	//	//doc.Find(".music .info p").Each(func(i int, s *goquery.Selection) {
	//	//	switch i % 2 {
	//	//	case 0:
	//	//		music.MusicName = s.Text()
	//	//	case 1:
	//	//		music.Composer = s.Text()
	//	//		musics = append(musics, music)
	//	//	}
	//	//})
	//
	//	for i := 0; i < len(musics); i++ {
	//		//infrastructure.InsertMusic(db, musics[len(musics)-i-1])
	//	}
	//
	//	time.Sleep(2 * time.Second)
	//
	//	//test用
	//	if i == 0 {
	//		break
	//	}
	//
	//}

}
