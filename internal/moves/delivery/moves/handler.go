package moves

import (
	"context"

	moves_grpc "github.com/RioRizkyRainey/pokedex/internal/moves/delivery/moves/grpc"
	"github.com/RioRizkyRainey/pokedex/internal/moves/usecase"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
	"google.golang.org/grpc"
)

func MovesServerGrpc(grpcServer grpc.ServiceRegistrar, movesUsecase usecase.MovesUsecase) {
	server := &server{
		usecase: movesUsecase,
	}
	moves_grpc.RegisterMovesHandlerServer(grpcServer, server)
}

type server struct {
	usecase usecase.MovesUsecase
	moves_grpc.UnimplementedMovesHandlerServer
}

func (s *server) GetMoves(ctx context.Context, params *moves_grpc.Params) (*moves_grpc.Data, error) {
	pokeName := params.Name
	moves, err := s.usecase.GetMoves(pokeName)

	if err != nil {
		result := &moves_grpc.Data{
			Status:  404,
			Message: "failed",
			Error: &moves_grpc.ErrorMessage{
				Message:        "Moves not found",
				Reason:         "not_found",
				ErrorUserTitle: err.Error(),
				ErrorUserMsg:   "Moves not found, please try another name",
			},
		}

		return result, nil
	}

	movesRpc := transformToRpc(moves)

	result := &moves_grpc.Data{
		Status:  200,
		Message: "success",
		Data:    movesRpc,
	}

	return result, nil
}

func transformToRpc(moves *model.Moves) *moves_grpc.Moves {
	return &moves_grpc.Moves{
		Id:             moves.PokID,
		Name:           moves.PokName,
		Weight:         moves.PokWeight,
		Height:         moves.PokHeight,
		BaseExperience: moves.PokBaseExperience,
	}
}
