package main

import (
	"context"
	"fmt"

	"github.com/yutaronakayama/building-http-clients-test/client"
)

func main() {
	client := client.NewClient()
	pokemon, err := client.GetPokemonByName(context.TODO(),"pikachu")
	if err != nil {
		panic(err)
	}
	fmt.Println(pokemon)
}