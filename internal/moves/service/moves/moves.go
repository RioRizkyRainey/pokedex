package moves

import (
	"database/sql"

	movesRpc "github.com/RioRizkyRainey/pokedex/internal/moves/delivery/moves"
	"github.com/RioRizkyRainey/pokedex/internal/moves/repository"
	"github.com/RioRizkyRainey/pokedex/internal/moves/usecase"
	"google.golang.org/grpc"
)

func Server(gserver *grpc.Server, connection *sql.DB) {
	movesRepo := repository.InitMovesRepository(connection)
	movesUsecase := usecase.MovesUsecase(movesRepo)

	movesRpc.MovesServerGrpc(gserver, movesUsecase)
}
