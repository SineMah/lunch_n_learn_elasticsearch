package build

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mincedmind.com/elasticsearch/elasticsearch"
	"mincedmind.com/elasticsearch/pokemon"
	"mincedmind.com/elasticsearch/pokemon/api"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

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

func getData() []pokemon.Pokemon {
	var pokemonList []pokemon.Entry
	directory, _ := os.Getwd()
	file, readErr := os.Open(fmt.Sprintf("%s/storage/pokedex.json", directory))

	if readErr != nil {
		fmt.Printf("File not found: %s\n", readErr)
		os.Exit(1)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	decodeErr := decoder.Decode(&pokemonList)

	if decodeErr != nil {
		fmt.Println("Error decoding JSON:", decodeErr)
		os.Exit(1)
	}

	return mapPokemonList(pokemonList)
}

func mapPokemonList(pokemonEntries []pokemon.Entry) []pokemon.Pokemon {
	var pokemonList []pokemon.Pokemon

	for _, entry := range pokemonEntries {
		var games []string
		var abilities []pokemon.Ability
		var moves []string
		var images []string
		meta, errMeta := getMeta(entry.ID)
		weight := 0
		height := 0

		if errMeta == nil {
			height = meta.Height
			weight = meta.Weight
			games = parseGames(meta)
			abilities = parseAbilities(meta)
			moves = parseMoves(meta)
			images = parseImages(meta)
		}

		entry := pokemon.Pokemon{
			ID:   entry.ID,
			Name: entry.Name.English,
			Type: entry.Type,
			Attributes: pokemon.Attributes{
				HP:             entry.Base.HP,
				Attack:         entry.Base.Attack,
				Defense:        entry.Base.Defense,
				SpecialAttack:  entry.Base.SpAttack,
				SpecialDefense: entry.Base.SpDefense,
				Speed:          entry.Base.Speed,
			},
			Abilities: abilities,
			Height:    height,
			Weight:    weight,
			Games:     games,
			Moves:     moves,
			Images:    images,
		}
		pokemonList = append(pokemonList, entry)
	}

	return pokemonList
}

func indexList(index string, pokemonList []pokemon.Pokemon) {
	namespace := uuid.MustParse("db11f1bb-8536-4c79-8c4c-08e28982fc1e")
	size := len(pokemonList)
	current := 1

	var wg sync.WaitGroup

	// index concurrently
	for _, entry := range pokemonList {
		wg.Add(1)

		go func(index string, entry pokemon.Pokemon) {
			defer wg.Done()

			id := uuid.NewSHA1(namespace, []byte(strconv.Itoa(entry.ID))).String()
			data, errMarshal := json.Marshal(entry)

			if errMarshal != nil {
				fmt.Printf("Error marshaling document: %s", errMarshal)
				os.Exit(1)
			}

			elasticsearch.AddDocument(index, id, data)

			fmt.Printf("%.2f%%\n", float64(current)*100/float64(size))

			current++
		}(index, entry)

		wg.Wait()
	}
}

func getMeta(id int) (api.Pokemon, error) {
	directory, _ := os.Getwd()
	pokemonApi := api.Pokemon{}

	fileBytes, errRead := os.ReadFile(fmt.Sprintf("%s/storage/meta/%d.json", directory, id))

	if errRead != nil {
		log.Printf("Error reading: %s", errRead)
		return pokemonApi, errRead
	}

	errDecode := json.Unmarshal(fileBytes, &pokemonApi)

	if errDecode != nil {
		log.Printf("Error decoding: %s", errDecode)
		return pokemonApi, errRead
	}

	return pokemonApi, nil
}

func parseGames(item api.Pokemon) []string {
	var games []string

	for _, game := range item.GameIndices {
		games = append(games, normalizeName(game.Version.Name))
	}

	return games
}

func parseMoves(item api.Pokemon) []string {
	var moves []string

	for _, move := range item.Moves {
		moves = append(moves, normalizeName(move.Move.Name))
	}

	return moves
}

func parseImages(item api.Pokemon) []string {
	var images []string

	images = append(images, item.Sprites.Other.DreamWorld.FrontDefault)
	images = append(images, item.Sprites.Other.Home.FrontDefault)
	images = append(images, item.Sprites.Other.OfficialArtwork.FrontDefault)

	return images
}

func parseAbilities(item api.Pokemon) []pokemon.Ability {
	var abilities []pokemon.Ability

	for _, ability := range item.Abilities {
		abilities = append(abilities, pokemon.Ability{
			Name:     normalizeName(ability.Ability.Name),
			Slot:     ability.Slot,
			IsHidden: ability.IsHidden,
		})
	}

	return abilities
}

func normalizeName(name string) string {
	return strings.ReplaceAll(name, "-", " ")
}
