package interfaces

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maooz4426/SDVX-Database/application/usecases"
	"net/http"
)

type Controller struct {
	*usecases.GetMusicUseCase
}

func NewController(u *usecases.GetMusicUseCase) *Controller {
	return &Controller{
		u,
	}
}

func (c *Controller) GetMusicData(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	vars := mux.Vars(r)
	key := vars["key"]

	musics, err := c.GetMusicUseCase.GetMusicData(ctx, key)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	js, err := json.MarshalIndent(musics, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(js)
}
