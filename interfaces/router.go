package interfaces

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maooz4426/SDVX-Database/application/usecases"
	"net/http"
)

func InitRouter(usc *usecases.GetMusicUseCase) {

	r := mux.NewRouter()

	c := Controller{usc}

	fmt.Println("Listening on port 8080")

	r.HandleFunc("/musics/{musicID}", c.GetMusicData)

	//これ最後にしないとhandleFuncが機能しない
	http.ListenAndServe(":8080", r)
}
