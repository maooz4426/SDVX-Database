package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maooz4426/SDVX-Database/domain"
	"github.com/maooz4426/SDVX-Database/infrastructure"
	"github.com/sclevine/agouti"
	"log"
	"strings"
	"time"
)

func main() {

	sqluser := "user"

	sqlpass := "password"

	sqldb := "sdvx_db"

	//dbのセットアップ
	//dsn := "user:password@tcp(127.0.0.1:3306)/sdvx_db?parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true&charset=utf8mb4", sqluser, sqlpass, sqldb)

	db, err := sql.Open("mysql", dsn)

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

	//テーブル作成
	infrastructure.CreateTable(db)

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

	selectBox := page.FindByName("page")
	err = selectBox.Select("2")

	if err != nil {
		log.Printf("Failed to select: %v", err)
	}

	time.Sleep(2 * time.Second)

	content, err := page.HTML()
	if err != nil {
		log.Printf("Failed to get html: %v", err)
	}

	reader := strings.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(reader)

	//type Music struct {
	//	name   string
	//	artist string
	//}

	var music domain.Music
	doc.Find(".music .info p").Each(func(i int, s *goquery.Selection) {

		switch i % 2 {
		case 0:
			music.Name = s.Text()
			//fmt.Println("Music:", music.Name)
		case 1:
			music.Artist = s.Text()
			//fmt.Println("Artist:", music.Artist)
			//fmt.Println(music)
			infrastructure.InsertMusic(db, music)
		}
	})

}
