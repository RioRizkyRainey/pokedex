package attack

import (
	"context"

	attack_grpc "github.com/RioRizkyRainey/pokedex/internal/attack/delivery/attack/grpc"
	"github.com/RioRizkyRainey/pokedex/internal/attack/usecase"
	"google.golang.org/grpc"
)

func AttackServerGrpc(grpcServer grpc.ServiceRegistrar, attackUsecase usecase.AttackUsecase) {
	server := &server{
		usecase: attackUsecase,
	}

	attack_grpc.RegisterAttackHandlerServer(grpcServer, server)
}

type server struct {
	usecase usecase.AttackUsecase
	attack_grpc.UnimplementedAttackHandlerServer
}

func (s *server) GetDamage(ctx context.Context, params *attack_grpc.Params) (*attack_grpc.Data, error) {
	damage, err := s.usecase.GetDamage(ctx, params.AttackName, params.AttackMove, params.DefendName)

	if err != nil {
		result := &attack_grpc.Data{
			Status:  404,
			Message: "failed",
			Error: &attack_grpc.ErrorMessage{
				Message:        "FAIL",
				Reason:         "not_found",
				ErrorUserTitle: err.Error(),
				ErrorUserMsg:   err.Error(),
			},
		}

		return result, nil
	}
	result := &attack_grpc.Data{
		Status:  200,
		Message: "success",
		Data:    *damage,
	}

	return result, nil
}
