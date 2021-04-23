package repository

import (
	"github.com/RioRizkyRainey/pokedex/internal/attack/delivery/moves"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type MovesRepository struct {
	Conn *moves.Client
}

type MovesRepositoryI interface {
	GetMovesByPokemon(name string) ([]*model.Moves, error)
}

func InitMovesRepository(Conn *moves.Client) *MovesRepository {
	return &MovesRepository{Conn: Conn}
}

func (r *MovesRepository) GetMovesByPokemon(name string) ([]*model.Moves, error) {
	return r.Conn.GetMoves(name)
}
