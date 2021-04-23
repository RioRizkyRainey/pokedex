package moves

import (
	"context"

	moves_grpc "github.com/RioRizkyRainey/pokedex/internal/moves/delivery/moves/grpc"
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

	moves = make([]*model.Moves, 0)

	for _, moveRpc := range data.Data {
		moves = append(moves, c.transformToModel(moveRpc))
	}

	return pokemon, nil
}

func transformToModel(pokemonRpc *pokemon_grpc.Pokemon) *model.Pokemon {
	return &model.Pokemon{
		PokID:             pokemonRpc.Id,
		PokName:           pokemonRpc.Name,
		PokHeight:         pokemonRpc.Height,
		PokWeight:         pokemonRpc.Weight,
		PokBaseExperience: pokemonRpc.BaseExperience,
	}
}
