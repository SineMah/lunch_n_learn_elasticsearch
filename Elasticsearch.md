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
```