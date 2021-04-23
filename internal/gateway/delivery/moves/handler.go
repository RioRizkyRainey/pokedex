package moves

import (
	"context"

	moves_grpc "github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/moves/grpc"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
	"google.golang.org/grpc"
)

func MovesClientGrpc(grpcClient *grpc.ClientConn) *Client {
	client := &Client{
		rpcClient: moves_grpc.NewMoveHandlerClient(grpcClient),
	}

	return client
}

type Client struct {
	rpcClient moves_grpc.MoveHandlerClient
}

func (c *Client) GetMoves(ctx context.Context, name string) ([]*model.Moves, error) {
	params := &moves_grpc.Params{
		Name: name,
	}

	data, err := c.rpcClient.GetMove(ctx, params)
	if err != nil {
		return nil, err
	}

	moves := make([]*model.Moves, 0)

	for _, moveRpc := range data.Data {
		moves = append(moves, transformToModel(moveRpc))
	}

	return moves, nil
}

func transformToModel(moveRpc *moves_grpc.Move) *model.Moves {
	return &model.Moves{
		MoveID:    moveRpc.Id,
		MovePower: &moveRpc.Power,
		MoveName:  moveRpc.Name,
		MovePp:    moveRpc.Pp,
	}
}
