package rest

import (
	"context"
	"encoding/json"
	"net/http"

	attackUsecase "github.com/RioRizkyRainey/pokedex/internal/gateway/usecase/attack"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	attUsecase attackUsecase.AttackUsecase
)

func InitializeRouter(router *mux.Router, a attackUsecase.AttackUsecase) {
	attUsecase = a
	router.HandleFunc("/api/attack", GetPokemon).Methods("GET")
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	attackName := params.Get("attack_name")
	attackMove := params.Get("attack_move")
	defendName := params.Get("defend_name")

	pokemon, err := attUsecase.GetDamage(context.Background(), attackName, attackMove, defendName)
	if err != nil {
		responsJson := &model.ResponseJSON{
			Message: err.Error(),
			Status:  200,
		}
		w.Header().Add("Content-Type", "application/json")
		bytes, err := json.Marshal(responsJson)
		if err != nil {
			log.Errorf("Can not marshal. Got %s", err)
		} else {
			i, err := w.Write(bytes)
			if err != nil {
				log.Errorf("Can not write byte stream. Got %s. %d bytes written", err, i)
			}
		}
		return
	}

	responsJson := &model.ResponseJSON{
		Message: "SUCCESS",
		Status:  200,
		Data:    pokemon,
	}

	w.Header().Add("Content-Type", "application/json")
	bytes, err := json.Marshal(responsJson)
	if err != nil {
		log.Errorf("Can not marshal. Got %s", err)
	} else {
		i, err := w.Write(bytes)
		if err != nil {
			log.Errorf("Can not write byte stream. Got %s. %d bytes written", err, i)
		}
	}
}
