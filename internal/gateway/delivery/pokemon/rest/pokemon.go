package rest

import (
	"context"
	"encoding/json"
	"net/http"

	pokeUsecase "github.com/RioRizkyRainey/pokedex/internal/gateway/usecase/pokemon"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	pokemonUsecase pokeUsecase.PokemonUsecase
)

func InitializeRouter(router *mux.Router, useCase pokeUsecase.PokemonUsecase) {
	pokemonUsecase = useCase
	router.HandleFunc("/api/pokemon/{name}", GetPokemon).Methods("GET")
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pokemonName := params["name"]

	pokemon, err := pokemonUsecase.GetPokemon(context.Background(), pokemonName)
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
