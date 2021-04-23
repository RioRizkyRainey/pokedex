package pokemon

import (
	"github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/pokemon"
	pokeRest "github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/pokemon/rest"
	pokemonRepo "github.com/RioRizkyRainey/pokedex/internal/gateway/repository/pokemon"
	pokeUsecase "github.com/RioRizkyRainey/pokedex/internal/gateway/usecase/pokemon"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func Server(router *mux.Router, pokemonClient *grpc.ClientConn) {
	pokeClient := pokemon.PokemonClientGrpc(pokemonClient)
	pokemonRepo := pokemonRepo.InitPokemonRepository(pokeClient)

	pokemonUsecase := pokeUsecase.InitPokemonUsecase(*pokemonRepo)

	pokeRest.InitializeRouter(router, pokemonUsecase)
}
