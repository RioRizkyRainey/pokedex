package pokemon

import (
	"database/sql"

	pokemonRpc "github.com/RioRizkyRainey/pokedex/internal/pokemon/delivery/pokemon"
	"github.com/RioRizkyRainey/pokedex/internal/pokemon/repository"
	"github.com/RioRizkyRainey/pokedex/internal/pokemon/usecase"
	"google.golang.org/grpc"
)

func Server(gserver *grpc.Server, connection *sql.DB) {
	pokemonRepo := repository.InitPokemonRepository(connection)
	pokemonUsecase := usecase.PokemonUsecase(pokemonRepo)

	pokemonRpc.PokemonServerGrpc(gserver, pokemonUsecase)
}
