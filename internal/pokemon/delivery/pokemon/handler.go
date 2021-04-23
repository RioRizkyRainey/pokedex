package pokemon

import (
	"context"

	pokemon_grpc "github.com/RioRizkyRainey/pokedex/internal/pokemon/delivery/pokemon/grpc"
	"github.com/RioRizkyRainey/pokedex/internal/pokemon/usecase"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
	"google.golang.org/grpc"
)

func PokemonServerGrpc(grpcServer grpc.ServiceRegistrar, pokemonUsecase usecase.PokemonUsecase) {
	server := &server{
		usecase: pokemonUsecase,
	}
	pokemon_grpc.RegisterPokemonHandlerServer(grpcServer, server)
}

type server struct {
	usecase usecase.PokemonUsecase
	pokemon_grpc.UnimplementedPokemonHandlerServer
}

func (s *server) GetPokemon(ctx context.Context, params *pokemon_grpc.Params) (*pokemon_grpc.Data, error) {
	pokeName := params.Name
	pokemon, err := s.usecase.GetPokemon(pokeName)

	if err != nil {
		result := &pokemon_grpc.Data{
			Status:  404,
			Message: "failed",
			Error: &pokemon_grpc.ErrorMessage{
				Message:        "Pokemon not found",
				Reason:         "not_found",
				ErrorUserTitle: err.Error(),
				ErrorUserMsg:   "Pokemon not found, please try another name",
			},
		}

		return result, nil
	}

	pokemonRpc := transformToRpc(pokemon)

	result := &pokemon_grpc.Data{
		Status:  200,
		Message: "success",
		Data:    pokemonRpc,
	}

	return result, nil
}

func transformToRpc(pokemon *model.Pokemon) *pokemon_grpc.Pokemon {
	return &pokemon_grpc.Pokemon{
		Id:             pokemon.PokID,
		Name:           pokemon.PokName,
		Weight:         pokemon.PokWeight,
		Height:         pokemon.PokHeight,
		BaseExperience: pokemon.PokBaseExperience,
		Attack:         pokemon.PokAttack,
		Defense:        pokemon.PokDefense,
	}
}
