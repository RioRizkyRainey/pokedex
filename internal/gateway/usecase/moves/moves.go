package moves

import (
	"context"

	movesRepo "github.com/RioRizkyRainey/pokedex/internal/gateway/repository/moves"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type MovesUsecase interface {
	GetMove(ctx context.Context, name string) ([]*model.Moves, error)
}

type usecase struct {
	movesRepository movesRepo.MovesRepository
}

func InitPokemonUsecase(movesRepository movesRepo.MovesRepository) MovesUsecase {
	return &usecase{
		movesRepository: movesRepository,
	}
}

func (u *usecase) GetMove(ctx context.Context, name string) ([]*model.Moves, error) {
	return u.movesRepository.GetMovesByPokemon(ctx, name)
}
