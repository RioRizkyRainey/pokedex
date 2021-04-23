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
	router.HandleFunc("/api/pokemon", GetPokemon).Methods("GET")
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pokemonName := query.Get("name")

	pokemon, err := pokemonUsecase.GetPokemon(context.Background(), pokemonName)
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
