package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RioRizkyRainey/pokedex/internal/pokemon/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemon(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"pok_id", "pok_name", "pok_height", "pok_weight", "pok_base_experience", "pok_attack", "pok_defense"}).
		AddRow(1, "bulbasaur", 12, 12, 3, 4, 6)

	query := regexp.QuoteMeta(`SELECT p.pok_id, p.pok_name, p.pok_height, p.pok_weight, p.pok_base_experience, b.b_atk as pok_attack, b.b_def as pok_defense
		FROM pokemon p
		INNER JOIN base_stats b ON b.pok_id = p.pok_id
		WHERE pok_name = ?`)

	mock.ExpectQuery(query).WillReturnRows(rows)

	pokemonRepo := repository.InitPokemonRepository(db)

	pokemon, err := pokemonRepo.GetPokemon("bulbasaur")

	assert.NoError(t, err)
	assert.NotNil(t, pokemon)
}
