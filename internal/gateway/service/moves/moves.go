package moves

import (
	"github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/moves"
	moveRest "github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/moves/rest"
	moveRepo "github.com/RioRizkyRainey/pokedex/internal/gateway/repository/moves"
	moveUsecase "github.com/RioRizkyRainey/pokedex/internal/gateway/usecase/moves"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func Server(router *mux.Router, moveClient *grpc.ClientConn) {
	movesClient := moves.MovesClientGrpc(moveClient)
	movesRepo := moveRepo.InitMovesRepository(movesClient)

	movesUsecase := moveUsecase.InitPokemonUsecase(*movesRepo)

	moveRest.InitializeRouter(router, movesUsecase)
}
