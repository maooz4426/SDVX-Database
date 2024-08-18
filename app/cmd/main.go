package main

import (
	"context"
	"github.com/maooz4426/SDVX-Database/repository"
	"github.com/maooz4426/SDVX-Database/services"
	"log"
	"os"
)

func main() {
	db, err := repository.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}
	//sqluser := "user"
	//
	//sqlpass := "password"
	//
	//sqldb := "sdvx_db"
	//
	////dbのセットアップ
	//
	//dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true&charset=utf8mb4", sqluser, sqlpass, sqldb)
	//
	//db, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i := 0; i < 6; i++ {
	//	if errcon := db.Ping(); errcon != nil {
	//		log.Println("Now Trying count:", i)
	//		log.Fatal("Failed to connect to database", errcon)
	//		db, _ = sql.Open("mysql", dsn)
	//	} else {
	//		break
	//	}
	//	time.Sleep(5 * time.Second)
	//}
	//
	//fmt.Println("Connected to database")

	rep := repository.NewMusicRepository(db)

	svc := services.NewService(rep)

	ctx := context.Background()

	switch os.Args[1] {
	case "register":
		svc.RegisterMusicData(ctx)
	}

}
