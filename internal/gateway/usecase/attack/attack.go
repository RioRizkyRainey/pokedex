package attack

import (
	"context"

	attackRepo "github.com/RioRizkyRainey/pokedex/internal/gateway/repository/attack"
)

type AttackUsecase interface {
	GetDamage(ctx context.Context, attackName string, attackMove string, defendName string) (int32, error)
}

type usecase struct {
	attRepository attackRepo.AttackRepository
}

func InitUsecase(attRepository attackRepo.AttackRepository) AttackUsecase {
	return &usecase{
		attRepository: attRepository,
	}
}

func (u *usecase) GetDamage(ctx context.Context, attackName string, attackMove string, defendName string) (int32, error) {
	return u.attRepository.GetDamage(ctx, attackName, attackMove, defendName)
}
