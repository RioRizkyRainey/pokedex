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
	GetMovesByPokemon(name string) ([]*model.Moves, error)
}

func InitMovesRepository(Conn *sql.DB) *MovesRepository {
	return &MovesRepository{Conn: Conn}
}

func (r *MovesRepository) GetMovesByPokemon(name string) ([]*model.Moves, error) {
	query := `SELECT moves.move_id, moves.move_name, moves.move_power, moves.move_pp, moves.move_accuracy, types.type_name as type
		FROM moves
		INNER JOIN pokemon_moves ON moves.move_id = pokemon_moves.move_id
		INNER JOIN pokemon ON pokemon.pok_id = pokemon_moves.pok_id
		INNER JOIN types ON moves.type_id = types.type_id
		WHERE pokemon.pok_name = ?`

	rows, err := r.Conn.Query(query, name)

	if err != nil {
		fmt.Println(err)
	}

	moves := make([]*model.Moves, 0)

	err = database.Scan(rows, &moves)

	if err != nil {
		return nil, err
	}

	return moves, nil
}
