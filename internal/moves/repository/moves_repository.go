package repository

import (
	"database/sql"
	"fmt"

	"github.com/RioRizkyRainey/pokedex/pkg/database"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type MovesRepository struct {
	Conn *sql.DB
}

type MovesRepositoryI interface {
	GetMoves(name string) model.Moves
}

func InitMovesRepository(Conn *sql.DB) *MovesRepository {
	return &MovesRepository{Conn: Conn}
}

func (r *MovesRepository) GetMoves(name string) (*model.Moves, error) {
	query := `SELECT *
		FROM moves
		WHERE pok_name = ?`

	rows, err := r.Conn.Query(query, name)

	if err != nil {
		fmt.Println(err)
	}

	moves := &model.Moves{}

	err = database.Scan(rows, moves)

	if err != nil {
		return nil, err
	}

	return moves, nil
}
