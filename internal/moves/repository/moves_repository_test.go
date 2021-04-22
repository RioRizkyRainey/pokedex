package repository_test

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
	rows := sqlmock.NewRows([]string{"pok_id", "pok_name", "pok_height", "pok_weight", "pok_base_experience"}).
		AddRow(1, "bulbasaur", 12, 12, 3)

	query := regexp.QuoteMeta(`SELECT *
						FROM moves
						WHERE pok_name = ?`)

	mock.ExpectQuery(query).WillReturnRows(rows)

	movesRepo := repository.InitMovesRepository(db)

	moves, err := movesRepo.GetMoves("bulbasaur")

	assert.NoError(t, err)
	assert.NotNil(t, moves)
}
