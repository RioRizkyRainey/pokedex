package pokemon

import (
	"context"

	"github.com/RioRizkyRainey/pokedex/internal/gateway/delivery/pokemon"
	"github.com/RioRizkyRainey/pokedex/pkg/model"
)

type PokemonRepository struct {
	Conn *pokemon.Client
}

type PokemonRepositoryI interface {
	GetPokemon(name string) (*model.Pokemon, error)
}

func InitPokemonRepository(Conn *pokemon.Client) *PokemonRepository {
	return &PokemonRepository{Conn: Conn}
}

func (r *PokemonRepository) GetPokemon(ctx context.Context, name string) (*model.Pokemon, error) {
	return r.Conn.GetPokemon(ctx, name)
}
