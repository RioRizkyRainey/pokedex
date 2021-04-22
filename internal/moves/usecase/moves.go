package usecase

import (
	"github.com/RioRizkyRainey/pokedex/internal/moves/repository"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type MovesUsecase interface {
	GetMovesByPokemon(name string) ([]*model.Moves, error)
}

type usecase struct {
	movesRepo *repository.MovesRepository
}

func InitMovesUsecase(movesRepo *repository.MovesRepository) MovesUsecase {
	return &usecase{
		movesRepo: movesRepo,
	}
}

func (u *usecase) GetMovesByPokemon(name string) ([]*model.Moves, error) {
	return u.movesRepo.GetMovesByPokemon(name)
}
