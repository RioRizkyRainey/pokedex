package usecase

import (
	"context"
	"errors"

	movesRepo "github.com/RioRizkyRainey/pokedex/internal/attack/repository/moves"
	pokemonRepo "github.com/RioRizkyRainey/pokedex/internal/attack/repository/pokemon"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type AttackUsecase interface {
	GetDamage(ctx context.Context, attackName string, attackMove, defendName string) (*int32, error)
}

type usecase struct {
	pokemonRepo pokemonRepo.PokemonRepository
	movesRepo   movesRepo.MovesRepository
}

func InitAttackUsecase(pokemonRepo pokemonRepo.PokemonRepository, movesRepo movesRepo.MovesRepository) AttackUsecase {
	return &usecase{
		pokemonRepo: pokemonRepo,
		movesRepo:   movesRepo,
	}
}

func (u *usecase) GetDamage(ctx context.Context, attackName string, attackMove, defendName string) (*int32, error) {
	attackPokemon, err := u.pokemonRepo.GetPokemon(ctx, attackName)

	if err != nil {
		return nil, err
	}

	defendPokemon, err := u.pokemonRepo.GetPokemon(ctx, defendName)

	if err != nil {
		return nil, err
	}

	attackMoves, err := u.movesRepo.GetMovesByPokemon(ctx, attackName)

	if err != nil {
		return nil, err
	}

	var isPokemonMove = false

	var move model.Moves

	for _, m := range attackMoves {
		if m.MoveName == attackMove {
			isPokemonMove = true
			move = *m
			break
		}
	}

	if !isPokemonMove {
		return nil, errors.New("Pokemon " + attackName + " doesn't have this move: " + attackMove)
	}

	movePower := *move.MovePower

	damage := ((2*50/5+2)*movePower*attackPokemon.PokAttack/defendPokemon.PokDefense)/50 + 2
	return &damage, nil
}
