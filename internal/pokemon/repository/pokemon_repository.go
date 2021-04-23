package repository

import (
	"database/sql"
	"fmt"

	"github.com/RioRizkyRainey/pokedex/pkg/database"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type PokemonRepository struct {
	Conn *sql.DB
}

type PokemonRepositoryI interface {
	GetPokemon(name string) (*model.Pokemon, error)
}

func InitPokemonRepository(Conn *sql.DB) *PokemonRepository {
	return &PokemonRepository{Conn: Conn}
}

func (r *PokemonRepository) GetPokemon(name string) (*model.Pokemon, error) {
	query := `SELECT p.pok_id, p.pok_name, p.pok_height, p.pok_weight, p.pok_base_experience, b.b_atk as pok_attack, b.b_def as pok_defense
		FROM pokemon p
		INNER JOIN base_stats b ON b.pok_id = p.pok_id
		WHERE pok_name = ?`

	rows, err := r.Conn.Query(query, name)

	if err != nil {
		fmt.Println(err)
	}

	pokemon := &model.Pokemon{}

	err = database.Scan(rows, pokemon)

	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
