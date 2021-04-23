package attack_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RioRizkyRainey/pokedex/internal/moves/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetMoves(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"move_id", "move_name", "move_power", "move_pp", "move_accuracy", "type"}).
		AddRow(1, "swords-dance", 12, 12, 3, "normal")

	query := regexp.QuoteMeta(`SELECT moves.move_id, moves.move_name, moves.move_power, moves.move_pp, moves.move_accuracy, types.type_name as type
		FROM moves
		INNER JOIN pokemon_moves ON moves.move_id = pokemon_moves.move_id
		INNER JOIN pokemon ON pokemon.pok_id = pokemon_moves.pok_id
		INNER JOIN types ON moves.type_id = types.type_id
		WHERE pokemon.pok_name = ?`)

	mock.ExpectQuery(query).WillReturnRows(rows)

	movesRepo := repository.InitMovesRepository(db)

	moves, err := movesRepo.GetMovesByPokemon("bulbasaur")

	assert.NoError(t, err)
	assert.NotNil(t, moves)
}
