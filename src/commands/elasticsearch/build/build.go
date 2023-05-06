package build

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	"log"
	"mincedmind.com/elasticsearch/elasticsearch"
	"os"
	"strconv"
	"time"
)

type PokemonEntry struct {
	ID   int `json:"id"`
	Name struct {
		English  string `json:"english"`
		Japanese string `json:"japanese"`
		Chinese  string `json:"chinese"`
		French   string `json:"french"`
	} `json:"name"`
	Type []string `json:"type"`
	Base struct {
		HP        int `json:"HP"`
		Attack    int `json:"Attack"`
		Defense   int `json:"Defense"`
		SpAttack  int `json:"Sp. Attack"`
		SpDefense int `json:"Sp. Defense"`
		Speed     int `json:"Speed"`
	} `json:"base"`
}

type Attributes struct {
	HP             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}

type Pokemon struct {
	ID         int
	Name       string
	Type       []string
	Attributes Attributes
}

func Start(params []string) {

	if len(params) == 0 {
		fmt.Println("No index defined")
		os.Exit(1)
	}

	alias := params[0]
	currentTime := time.Now()
	currentIndex := fmt.Sprintf("search_%s_%d", alias, currentTime.Unix())

	fmt.Println(alias, currentIndex)

	elasticsearch.FlushIndices(alias)
	elasticsearch.CreateIndex(currentIndex)

	indexList(currentIndex, getData())

	elasticsearch.AddAlias(currentIndex, alias)
}

func getData() []Pokemon {
	var pokemons []PokemonEntry
	directory, _ := os.Getwd()
	file, readErr := os.Open(fmt.Sprintf("%s/storage/pokedex.json", directory))

	if readErr != nil {
		fmt.Printf("File not found: %s\n", readErr)
		os.Exit(1)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	decodeErr := decoder.Decode(&pokemons)

	if decodeErr != nil {
		fmt.Println("Error decoding JSON:", decodeErr)
		os.Exit(1)
	}

	return mapPokemonList(pokemons)
}

func mapPokemonList(pokemonEntries []PokemonEntry) []Pokemon {
	var pokemonList []Pokemon

	for _, entry := range pokemonEntries {
		pokemon := Pokemon{
			ID:   entry.ID,
			Name: entry.Name.English,
			Type: entry.Type,
			Attributes: Attributes{
				HP:             entry.Base.HP,
				Attack:         entry.Base.Attack,
				Defense:        entry.Base.Defense,
				SpecialAttack:  entry.Base.SpAttack,
				SpecialDefense: entry.Base.SpDefense,
				Speed:          entry.Base.Speed,
			},
		}
		pokemonList = append(pokemonList, pokemon)
	}

	return pokemonList
}

func indexList(index string, pokemonList []Pokemon) {
	client := elasticsearch.GetClient()
	namespace := uuid.MustParse("db11f1bb-8536-4c79-8c4c-08e28982fc1e")

	for _, pokemon := range pokemonList {
		data, errMarshal := json.Marshal(pokemon)
		if errMarshal != nil {
			log.Fatalf("Error marshaling document: %s", errMarshal)
		}

		req := esapi.IndexRequest{
			Index:      index,
			DocumentID: uuid.NewSHA1(namespace, []byte(strconv.Itoa(pokemon.ID))).String(),
			Body:       bytes.NewReader(data),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), client)

		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}

		defer res.Body.Close()

		if res.IsError() {
			fmt.Printf("Error: %s\n", res.String())
		}
	}
}
