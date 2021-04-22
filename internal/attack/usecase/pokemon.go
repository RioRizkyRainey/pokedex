package usecase

import (
	"github.com/RioRizkyRainey/pokedex/internal/pokemon/repository"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type PokemonUsecase interface {
	GetPokemon(name string) (*model.Pokemon, error)
}

type usecase struct {
	pokemonRepo repository.PokemonRepository
}

func InitPokemonUsecase(pokemonRepo repository.PokemonRepository) PokemonUsecase {
	return &usecase{
		pokemonRepo: pokemonRepo,
	}
}

func (u *usecase) GetPokemon(name string) (*model.Pokemon, error) {
	return u.pokemonRepo.GetPokemon(name)
}
