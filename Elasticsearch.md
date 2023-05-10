# Elasticsearch

## Queries
```json
GET _cat/indices
GEt _cat/aliases
GET pokemon_index/_mapping

GET pokemon_index/_search
{
  "query": {
    "match_all": {}
  }
}

GET pokemon_index/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "wildcard": {
            "Name": {
              "value": "Ch*d",
              "boost": 1,
              "rewrite": "constant_score"
            }
          }
        }
      ]
    }
  }
}

GET pokemon_index/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "Attributes.Attack": {
              "gte": 120,
              "lte": 140
            }
          }
        }
      ]
    }
  }
}

GET pokemon_index/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "Type": {
              "value": "Steel",
              "boost": 1
            }
          }
        }
      ]
    }
  }
}

GET pokemon_index/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "term": {
            "Type": {
              "value": "Steel",
              "boost": 1
            }
          }
        }
      ],
      "minimum_should_match": 1
    }
  }
}

GET pokemon_index/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "Type": {
              "value": "Steel",
              "boost": 1
            }
          }
        },
        {
          "term": {
            "Type": {
              "value": "Rock",
              "boost": 1
            }
          }
        }
      ]
    }
  }
}

GET pokemon_index/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "term": {
            "Type": {
              "value": "Steel",
              "boost": 1
            }
          }
        },
        {
          "term": {
            "Type": {
              "value": "Rock",
              "boost": 1
            }
          }
        }
      ],
      "minimum_should_match": 1
    }
  }
}

GET pokemon_index/_search
{
  "aggs": {
    "games_versions": {
      "terms": {
        "field": "Games",
        "size": 10000,
        "order": { "_count": "asc" }
      }
    }
  },
  "size": 0
}

GET pokemon_index/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "nested": {
            "path": "Abilities",
            "query": {
              "bool": {
                "must": [
                  {
                    "match": {
                      "Abilities.Name": {
                        "query": "flash fire",
                        "operator": "and"
                      }
                    }
                  }
                ]
              }
            }
          }
        },
        {
          "nested": {
            "path": "Abilities",
            "query": {
              "bool": {
                "must": [
                  {
                    "match": {
                      "Abilities.Name": {
                        "query": "justified",
                        "operator": "and"
                      }
                    }
                  }
                ]
              }
            }
          }
        }
      ],
      "minimum_should_match": 1
    }
  }
}
```

### Places-Query
```JSON
{
    "query": {
        "bool": {
            "must": [
                {
                    "nested": {
                        "path": "attributes",
                        "query": {
                            "bool": {
                                "must": [
                                    {
                                        "match": {
                                            "attributes.value.name": {
                                                "query": "vendor",
                                                "operator": "and",
                                                "auto_generate_synonyms_phrase_query": false,
                                                "fuzziness": "0"
                                            }
                                        }
                                    },
                                    {
                                        "match": {
                                            "attributes.path": {
                                                "query": "place_type",
                                                "operator": "and",
                                                "auto_generate_synonyms_phrase_query": false,
                                                "fuzziness": "0"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    }
                },
                {
                    "nested": {
                        "path": "attributes",
                        "query": {
                            "bool": {
                                "must": [
                                    {
                                        "match": {
                                            "attributes.value.code": {
                                                "query": "DE",
                                                "operator": "and",
                                                "auto_generate_synonyms_phrase_query": false,
                                                "fuzziness": "0"
                                            }
                                        }
                                    },
                                    {
                                        "match": {
                                            "attributes.path": {
                                                "query": "country",
                                                "operator": "and",
                                                "auto_generate_synonyms_phrase_query": false,
                                                "fuzziness": "0"
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    }
                },
                {
                    "nested": {
                        "path": "attributes",
                        "query": {
                            "bool": {
                                "should": [
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "7659ecd6-eccc-5393-95a2-a2728ba31011",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "963f9565-7459-5f30-b05b-dc5a14fc68a7",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "fd2f19c2-b6d7-54a3-a1a7-bcdc4996225f",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "ecc6613b-b128-5f9b-8ecb-ea5bbb767f0b",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "c7e4c22f-320e-5d2f-8f2f-ad1ad047eaba",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "b45b2745-d2d9-5bd7-862e-9e85a2812702",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "d0eb9cbd-5ed8-505e-b3d2-8bd467e8ee95",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "0d23e015-4dc9-5323-b702-e0430237e4f0",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "ede586f0-18cb-5553-86a5-f9d0a2e79d8e",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "6bfe6168-d8ab-5a05-b4d6-c82d00eb6bd7",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "522386bb-e169-5749-a463-819397ab8ab2",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "f9bdba1f-52c6-5934-88bb-7acc4b5f0ba5",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "20456e72-043a-516c-8395-f7bc90bd01bb",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "58b68f36-e07f-5d0b-bb44-72065862cbb8",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "7ce1a7d5-f166-5cda-82d7-c9a42f9f0d05",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "9989aabe-b4bf-50cd-8e23-3919c845f48c",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "0a406023-55c7-594f-9c10-fcf8d48d1ded",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "f11172af-efba-5ca7-a0bd-7c88886b4d6c",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "3fb1b7b1-3954-57c1-a348-4a61ea7121c8",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "9924daff-c45f-5304-939c-860c7a39effc",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "ee84d483-d657-5b45-9fc8-386c8b142b3e",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "5611b5a4-a2bf-4041-86f7-4eb64e3bd581",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class_group",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "match": {
                                                        "attributes.value.uuid": {
                                                            "query": "f4a84681-8f50-48c2-8b34-ec2d1720f737",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                },
                                                {
                                                    "match": {
                                                        "attributes.path": {
                                                            "query": "device_class_group",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                ],
                                "minimum_should_match": 1
                            }
                        }
                    }
                },
                {
                    "nested": {
                        "path": "agent",
                        "query": {
                            "bool": {
                                "should": [
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "bool": {
                                                        "must": [
                                                            {
                                                                "range": {
                                                                    "agent.production_count": {
                                                                        "boost": 1,
                                                                        "gte": 12
                                                                    }
                                                                }
                                                            }
                                                        ]
                                                    }
                                                }
                                            ],
                                            "should": [
                                                {
                                                    "match": {
                                                        "agent.partnership_agreements.name": {
                                                            "query": "WG PVA",
                                                            "operator": "and",
                                                            "auto_generate_synonyms_phrase_query": false,
                                                            "fuzziness": "0"
                                                        }
                                                    }
                                                }
                                            ],
                                            "must_not": [],
                                            "minimum_should_match": 1
                                        }
                                    },
                                    {
                                        "bool": {
                                            "must": [
                                                {
                                                    "term": {
                                                        "agent.available.exists": false
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                ],
                                "minimum_should_match": 1
                            }
                        }
                    }
                }
            ],
            "filter": [
                {
                    "geo_distance": {
                        "distance": "2000km",
                        "location": [
                            16.3739206,
                            48.2082647
                        ]
                    }
                }
            ]
        }
    },
    "from": 0,
    "size": 5,
    "_source": true,
    "script_fields": {
        "distance": {
            "script": {
                "lang": "painless",
                "source": "double distance = doc['location'].arcDistance(params.point.latitude, params.point.longitude);
int dividend = 1;

if(params.distance_unit == 'km') {
    dividend = 1000;
}

return distance\/dividend;
",
                "params": {
                    "distance_unit": "m",
                    "point": {
                        "latitude": 48.2082647,
                        "longitude": 16.3739206
                    }
                }
            }
        }
    },
    "sort": [
        {
            "_geo_distance": {
                "location": [
                    16.3739206,
                    48.2082647
                ],
                "order": "asc",
                "unit": "km",
                "mode": "min",
                "distance_type": "arc",
                "ignore_unmapped": true
            }
        }
    ]
}
```