package pokemon

import (
	"context"

	pokemon_grpc "github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/pokemon/grpc"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
	"google.golang.org/grpc"
)

func PokemonClientGrpc(grpcClient *grpc.ClientConn) *Client {
	client := &Client{
		rpcClient: pokemon_grpc.NewPokemonHandlerClient(grpcClient),
	}

	return client
}

type Client struct {
	rpcClient pokemon_grpc.PokemonHandlerClient
}

func (c *Client) GetPokemon(ctx context.Context, name string) (*model.Pokemon, error) {
	params := &pokemon_grpc.Params{
		Name: name,
	}

	data, err := c.rpcClient.GetPokemon(ctx, params)

	if err != nil {
		return nil, err
	}

	pokemon := transformToModel(data.Data)
	return pokemon, nil
}

func transformToModel(pokemonRpc *pokemon_grpc.Pokemon) *model.Pokemon {
	return &model.Pokemon{
		PokID:             pokemonRpc.Id,
		PokName:           pokemonRpc.Name,
		PokHeight:         pokemonRpc.Height,
		PokWeight:         pokemonRpc.Weight,
		PokBaseExperience: pokemonRpc.BaseExperience,
		PokAttack:         pokemonRpc.Attack,
		PokDefense:        pokemonRpc.Defense,
	}
}
