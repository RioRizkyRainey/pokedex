package pokemon

import (
	"context"

	pokemonRepo "github.com/RioRizkyRainey/pokedex/internal/gateway/repository/pokemon"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type PokemonUsecase interface {
	GetPokemon(ctx context.Context, name string) (*model.Pokemon, error)
}

type usecase struct {
	pokemonRepo pokemonRepo.PokemonRepository
}

func InitPokemonUsecase(pokemonRepo pokemonRepo.PokemonRepository) PokemonUsecase {
	return &usecase{
		pokemonRepo: pokemonRepo,
	}
}

func (u *usecase) GetPokemon(ctx context.Context, name string) (*model.Pokemon, error) {
	return u.pokemonRepo.GetPokemon(ctx, name)
}
