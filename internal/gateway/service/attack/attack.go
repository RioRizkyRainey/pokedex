package attack

import (
	"github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/attack"
	attackRest "github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/attack/rest"
	attackRepo "github.com/RioRizkyRainey/pokedex/internal/gateway/repository/attack"
	attackUsecase "github.com/RioRizkyRainey/pokedex/internal/gateway/usecase/attack"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func Server(router *mux.Router, attackClient *grpc.ClientConn) {
	attClient := attack.AttackClientGrpc(attackClient)
	attRepo := attackRepo.InitAttackRepository(attClient)

	attUsecase := attackUsecase.InitUsecase(*attRepo)

	attackRest.InitializeRouter(router, attUsecase)
}
