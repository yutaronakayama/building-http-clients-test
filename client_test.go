package client

import (
	"context"
	"testing"

	_ "github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Run("GetPokemonByName", func(t *testing.T) {
		myclient := NewClient()
		pokemon, err := myclient.GetPokemonByName(context.Background(), "pikachu")
		//fmt.Println(pokemon)
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", pokemon.Name)
	})
}
