package database_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RioRizkyRainey/pokedex/pkg/database"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestScanSlice(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"pok_id", "pok_name", "pok_height", "pok_weight", "pok_base_experience"}).
		AddRow(1, "name-1", 12, 12, 3).
		AddRow(2, "name-2", 13, 12, 2).
		AddRow(3, "name-3", 14, 12, 1)

	query := regexp.QuoteMeta(`SELECT pok_id, pok_name, pok_height, pok_weight, pok_base_experience
			FROM pokemon`)

	mock.ExpectQuery(query).WillReturnRows(rows)

	dbQuery := `SELECT pok_id, pok_name, pok_height, pok_weight, pok_base_experience
	FROM pokemon`

	dbRows, err := db.Query(dbQuery)

	pokemons := make([]*model.Pokemon, 0)
	database.Scan(dbRows, &pokemons)

	assert.NoError(t, err)
	assert.Len(t, pokemons, 3)
}

func TestScanStruct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"pok_id", "pok_name", "pok_height", "pok_weight", "pok_base_experience"}).
		AddRow(1, "name-1", 12, 12, 3)

	query := regexp.QuoteMeta(`SELECT pok_id, pok_name, pok_height, pok_weight, pok_base_experience
	FROM pokemon`)

	mock.ExpectQuery(query).WillReturnRows(rows)

	dbQuery := `SELECT pok_id, pok_name, pok_height, pok_weight, pok_base_experience
	FROM pokemon`

	dbRows, err := db.Query(dbQuery)

	pokemon := &model.Pokemon{}
	database.Scan(dbRows, &pokemon)

	assert.NoError(t, err)
	assert.NotNil(t, pokemon.PokID)
	assert.NotNil(t, pokemon.PokName)
	assert.NotNil(t, pokemon.PokHeight)
}

func TestToSnakeCase(t *testing.T) {
	name := "ImageURL"
	result := database.ToSnakeCase(name)
	assert.Equal(t, result, "image_url")
}
