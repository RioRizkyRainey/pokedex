package usecase

import "github.com/RioRizkyRainey/pokedex/pkg/model"

type AttackUsecase interface {
	GetDamage(attackName string, attackMove, defendName string) (*int32, error)
}

type usecase struct {
	pokemonRepo pokemonRepo.PokemonRepository
	movesRepo   movesRepo.MovesRepository
}

func InitAttackUsecase(pokemonRepo pokemonRepo.PokemonRepository, movesRepo movesRepo.MovesRepository) PokemonUsecase {
	return &usecase{
		pokemonRepo: pokemonRepo,
		movesRepo:   movesRepo,
	}
}

func (u *usecase) GetDamage(attackName string, attackMove, defendName string) (*int32, error) {
	attackPokemon, err := u.pokemonRepo.GetPokemon(attackName)

	if err != nil {
		return nil, err
	}

	defendPokemon, err := u.pokemonRepo.GetPokemon(defendName)

	if err != nil {
		return nil, err
	}

	attackMoves, err := u.movesRepo.GetMovesByPokemon(attackName)

	if err != nil {
		return nil, err
	}

	var isPokemonMove = false

	var move model.Moves

	for _, m := range moves {
		if m.Name == attackMove {
			isPokemonMove = true
			move = m
			break
		}
	}

	if !isPokemonMove {
		return nil.errors.New("Pokemon " + attackPokemon + " doesn't have this move: " + attackMove)
	}

	damage := (2*50/5 + 2) * move
}
