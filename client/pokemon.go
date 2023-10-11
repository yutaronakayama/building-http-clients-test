package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var pokeAPIEndpoint string = "https://pokeapi.co/api/v2/pokemon/"

type MyIntTransport int

func (t *MyIntTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	return nil, nil
}

type MyTransport struct {
	rtp         http.RoundTripper
	elapsedTime time.Duration
}

// (*MyTransport)にキャストして、http.RoundTripperとして代入可能かどうかを見る
var _ http.RoundTripper = (*MyTransport)(nil)

// おかわりタスク：httpClient cli.Doを実行するとリクエストを送ってレスポンスを受け取る
func (t *MyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	end := time.Now()
	t.elapsedTime = end.Sub(start)
	return res, err
}

func (c *Client) GetPokemonByName(ctx context.Context, pokemonName string) (Pokemon, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		pokeAPIEndpoint+pokemonName,
		nil,
	)
	if err != nil {
		return Pokemon{}, err
	}

	//リクエストヘッダーを設定
	req.Header.Add("Accept", "application/json")

	//RoundTripperの設定
	tp := &MyTransport{}
	c.httpClient.Transport = tp

	//http.Get()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	fmt.Println(tp.elapsedTime)
	
	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code returned from the pokeapi")
	}

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	//t.RoundTrip(req)メソッドを使う場合
	// //t := &Transport{}
	// res, err := tp.RoundTrip(req)
	// if err != nil {
	// 	return Pokemon{}, err
	// }
	// var pokemonRoundTripper Pokemon
	// err = json.NewDecoder(res.Body).Decode(&pokemonRoundTripper)
	// if err != nil {
	// 	return Pokemon{}, err
	// }
	// defer res.Body.Close()

	return pokemon, nil
}
