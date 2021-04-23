package attack

import (
	"context"

	"github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/attack"
)

type AttackRepository struct {
	Conn *attack.Client
}

type AttackRepositoryI interface {
	GetDamage(attackName string, attackMove string, defendName string) (int32, error)
}

func InitMovesRepository(Conn *attack.Client) *AttackRepository {
	return &AttackRepository{Conn: Conn}
}

func (r *AttackRepository) GetDamage(ctx context.Context, attackName string, attackMove string, defendName string) (int32, error) {
	return r.Conn.GetDamage(ctx, attackName, attackMove, defendName)
}
