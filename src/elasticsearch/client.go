package elasticsearch

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"os"
	"strings"
	"sync"
)

type client struct {
	instance *elasticsearch.Client
}

var elasticInstance *client
var once sync.Once

func initClient() *elasticsearch.Client {
	config := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_ADDRESS"),
		},
	}

	es, err := elasticsearch.NewClient(config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return es
}

func GetClient() *elasticsearch.Client {

	once.Do(func() {
		newClient := initClient()
		elasticInstance = &client{newClient}
	})

	return elasticInstance.instance
}

func CreateIndex(name string) {
	client := GetClient()
	body := `{
        "settings": {
            "number_of_shards": 1,
            "number_of_replicas": 0
        },
        "mappings": {
            "properties": {
                "Name": {
                    "type": "keyword"
                },
				"Type": {
                    "type": "keyword"
                },
                "ID": {
                    "type": "integer"
                },
				"Abilities": {
					"type": "nested"				
				},
				"Attributes": {
					"properties": {
						"Attack": {
							"type": "integer"
						},
						"Defense": {
							"type": "integer"
						},
						"HP": {
							"type": "integer"
						},
						"SpecialAttack": {
							"type": "integer"
						},
						"SpecialDefense": {
							"type": "integer"
						},
						"Speed": {
							"type": "integer"
						}
					}
				}
            }
        }
    }`

	req := esapi.IndicesCreateRequest{
		Index: name,
		Body:  bytes.NewReader([]byte(body)),
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Printf("Error creating the index: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Printf("Error creating the index: %s\n", res.String())
		os.Exit(1)
	}
}

func AddDocument(index string, id string, data []byte) {
	client := GetClient()
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), client)

	if err != nil {
		fmt.Printf("Error getting response: %s", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.IsError() {
		fmt.Printf("Error: %s\n", res.String())
		os.Exit(1)
	}
}

func AddAlias(index string, alias string) {
	client := GetClient()
	body := fmt.Sprintf(`{
            "actions": [
                {
                    "add": {
                        "index": "%s",
                        "alias": "%s"
                    }
                }
            ]
        }`, index, alias)
	req := esapi.IndicesUpdateAliasesRequest{
		Body: bytes.NewReader([]byte(body)),
	}

	_, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Println("Error adding alias:", err)
		os.Exit(1)
	}
}

func GetIndices(alias string) []string {
	var indices []string
	client := GetClient()
	pattern := fmt.Sprintf("search_%s_*", alias)

	req := esapi.CatIndicesRequest{
		Index: []string{pattern},
	}

	res, errIndices := req.Do(context.Background(), client)

	if errIndices != nil {
		fmt.Println("Error listing indices:", errIndices)
		os.Exit(1)
	}

	if res.IsError() {
		fmt.Printf("Error listing indices: %s\n", res.String())
		os.Exit(1)
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(fields) > 2 {
			indices = append(indices, fields[2])
		}
	}

	return indices
}

func FlushIndices(pattern string) {
	indices := GetIndices(pattern)

	for _, index := range indices {
		flushIndex(index)
	}
}

func flushIndex(index string) {
	client := GetClient()

	req := esapi.IndicesDeleteRequest{
		Index: []string{index},
	}

	res, err := req.Do(context.Background(), client)

	if err != nil {
		fmt.Println("Error deleting index:", err)
		os.Exit(1)
	}

	if res.IsError() {
		fmt.Printf("Error deleting index: %s", res.String())
		os.Exit(1)
	}
}
