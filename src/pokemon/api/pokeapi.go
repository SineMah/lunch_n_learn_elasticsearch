package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	GameIndices []struct {
		Version struct {
			Name string `json:"name"`
		} `json:"version"`
	} `json:"game_indices"`
	Height int `json:"height"`
	ID     int `json:"id"`
	Moves  []struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
	} `json:"moves"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
	Sprites struct {
		Other struct {
			DreamWorld struct {
				FrontDefault string  `json:"front_default"`
				FrontFemale  *string `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string  `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       string  `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
		} `json:"other"`
	} `json:"sprites"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func FetchAll() {
	var wg sync.WaitGroup

	for i := 1; i <= 1010; i += 1 {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			log.Printf("Saving pokemon %d", id)

			Fetch(id)
		}(i)

		wg.Wait()
	}
}

func Fetch(id int) {
	var pokemon Pokemon
	resp, errRequest := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id))

	if errRequest != nil {
		fmt.Println("Request error: ", errRequest)
		return
	}

	defer resp.Body.Close()

	errDecode := json.NewDecoder(resp.Body).Decode(&pokemon)

	if errDecode != nil {
		fmt.Println("Decode error: ", errDecode)
		return
	}

	directory, _ := os.Getwd()
	file, errCreate := os.Create(fmt.Sprintf("%s/storage/meta/%d.json", directory, id))

	if errCreate != nil {
		fmt.Println("File create error: ", errCreate)
		return
	}

	defer file.Close()

	errStore := json.NewEncoder(file).Encode(pokemon)

	if errStore != nil {
		fmt.Println("File store error: ", errStore)
		return
	}
}
