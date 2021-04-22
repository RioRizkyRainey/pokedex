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
	moves_grpc.RegisterMoveHandlerServer(grpcServer, server)
}

type server struct {
	usecase usecase.MovesUsecase
	moves_grpc.UnimplementedMoveHandlerServer
}

func (s *server) GetMove(ctx context.Context, params *moves_grpc.Params) (*moves_grpc.Data, error) {
	pokeName := params.Name
	moves, err := s.usecase.GetMovesByPokemon(pokeName)

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

	movesRpc := make([]*moves_grpc.Move, 0)
	for _, move := range moves {
		movesRpc = append(movesRpc, transformToRpc(move))
	}

	result := &moves_grpc.Data{
		Status:  200,
		Message: "success",
		Data:    movesRpc,
	}

	return result, nil
}

func transformToRpc(moves *model.Moves) *moves_grpc.Move {
	return &moves_grpc.Move{
		Id:       moves.MoveID,
		Name:     moves.MoveName,
		Power:    *moves.MovePower,
		Pp:       moves.MovePp,
		Accuracy: *moves.MoveAccuracy,
	}
}
