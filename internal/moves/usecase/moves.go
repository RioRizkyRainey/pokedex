package usecase

import (
	"github.com/RioRizkyRainey/pokedex/internal/moves/repository"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type MovesUsecase interface {
	GetMoves(name string) (*model.Moves, error)
}

type usecase struct {
	movesRepo repository.MovesRepository
}

func InitMovesUsecase(movesRepo repository.MovesRepository) MovesUsecase {
	return &usecase{
		movesRepo: movesRepo,
	}
}

func (u *usecase) GetMoves(name string) (*model.Moves, error) {
	return u.movesRepo.GetMoves(name)
}
