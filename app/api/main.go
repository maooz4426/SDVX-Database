package main

import (
	"github.com/maooz4426/SDVX-Database/application/usecases"
	"github.com/maooz4426/SDVX-Database/infrastructure"
	"github.com/maooz4426/SDVX-Database/interfaces"
	"log"
)

func main() {
	db, err := infrastructure.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}

	rep := infrastructure.NewMusicRepository(db)

	usc := usecases.NewGetMusicUseCase(rep)

	interfaces.InitRouter(usc)

}
