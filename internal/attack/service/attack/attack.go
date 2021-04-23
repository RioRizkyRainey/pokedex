package pokemon

import (
	attackRpc "github.com/RioRizkyRainey/pokedex/internal/attack/delivery/attack"
	pokemonRepo "github.com/RioRizkyRainey/pokedex/internal/attack/repository/pokemon"
	"github.com/RioRizkyRainey/pokedex/internal/attack/usecase"
	"google.golang.org/grpc"
)

func Server(gserver *grpc.Server, pokemonClient *grpc.ClientConnInterface, moveClient *grpc.ClientConnInterface) {
	pokemonRepo := pokemonRepo.InitPokemonRepository(pokemonClient)
	attackUsecase := usecase.AttackUsecase()

	attackRpc.PokemonServerGrpc(gserver, pokemonUsecase)
}
