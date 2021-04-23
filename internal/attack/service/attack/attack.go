package attack

import (
	attackRpc "github.com/RioRizkyRainey/pokedex/internal/attack/delivery/attack"
	"github.com/RioRizkyRainey/pokedex/internal/attack/delivery/moves"
	"github.com/RioRizkyRainey/pokedex/internal/attack/delivery/pokemon"
	movesRepo "github.com/RioRizkyRainey/pokedex/internal/attack/repository/moves"
	pokemonRepo "github.com/RioRizkyRainey/pokedex/internal/attack/repository/pokemon"
	"github.com/RioRizkyRainey/pokedex/internal/attack/usecase"
	"google.golang.org/grpc"
)

func Server(gserver *grpc.Server, pokemonClient *grpc.ClientConn, moveClient *grpc.ClientConn) {
	pokeClient := pokemon.PokemonClientGrpc(pokemonClient)
	pokemonRepo := pokemonRepo.InitPokemonRepository(pokeClient)

	movesClient := moves.MovesClientGrpc(moveClient)
	movesRepo := movesRepo.InitMovesRepository(movesClient)

	attackUsecase := usecase.InitAttackUsecase(*pokemonRepo, *movesRepo)

	attackRpc.AttackServerGrpc(gserver, attackUsecase)
}
