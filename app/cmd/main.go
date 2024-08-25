package main

import (
	"context"
	"github.com/maooz4426/SDVX-Database/application/usecases"
	"github.com/maooz4426/SDVX-Database/infrastructure"
	"log"
	"os"
)

func main() {
	db, err := infrastructure.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}

	rep := infrastructure.NewMusicRepository(db)

	usc := usecases.NewRegisterMusicData(rep)

	ctx := context.Background()

	switch os.Args[1] {
	case "register":
		usc.Register(ctx)
	}

}
