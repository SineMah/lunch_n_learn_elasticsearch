package routes

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"mincedmind.com/elasticsearch/elasticsearch"
	"sync"
)

var once sync.Once
var cache map[string][]Pokemon

func Search(ctx *fiber.Ctx) error {
	var buf bytes.Buffer
	index := ctx.Params("index")
	body := string(ctx.Body())
	hash := createMd5(index + body)

	//getMap()

	_, ok := cache[hash]

	if ok == false {
		buf.WriteString(body)

		cache[hash] = searchElasticsearch(ctx, index, buf)
	}

	return ctx.JSON(cache[hash])
}

func Hello(ctx *fiber.Ctx) error {
	version := Version{
		App:     "learn ya know",
		Version: "1.0.0",
	}
	return ctx.JSON(version)
}

func searchElasticsearch(ctx *fiber.Ctx, index string, body bytes.Buffer) []Pokemon {
	var data []Pokemon
	client := elasticsearch.GetClient()

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		client.Search.WithBody(&body),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	var response map[string]interface{}
	errDecode := json.NewDecoder(res.Body).Decode(&response)

	if errDecode != nil {
		return data
	}

	if res.IsError() {
		fmt.Printf("[%s] %s: %s\n",
			res.Status(),
			response["error"].(map[string]interface{})["type"],
			response["error"].(map[string]interface{})["reason"],
		)

		return data
	}

	for _, hit := range response["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		var pokemon Pokemon

		pokemonJson, errMarshal := json.Marshal(source)

		if errMarshal == nil {

			errUnmarshal := json.Unmarshal(pokemonJson, &pokemon)

			if errUnmarshal == nil {
				data = append(data, pokemon)
			}
		}
	}

	return data
}

func createMd5(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)

	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

func InitMap() map[string][]Pokemon {

	once.Do(func() {
		cache = make(map[string][]Pokemon)
	})

	return cache
}
