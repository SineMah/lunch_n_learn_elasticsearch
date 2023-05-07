# Elasticsearch Example
Setup-Beschreibung

Weiterführende Informationen befinden sich in der Datei `Elasticsearch.md`.
## Requirements
- Docker
- Browser

## Installation
Docker-Container liefern alle Abhängigkeiten
- Elasticsearch
- Kibana
- Go 1.20.4
```bash
docker-compose build
```
Container starten
```bash
docker-compose up
```
Interaktive Shell des Containers go_app öffnen
```bash
docker exec -it go_app sh
```
Im go_app Container die env-Datei umkopieren
```bash
cp .env.example .env
```
Go-Dependencies installieren, Binary bauen, CLI-Command build ausführen, um einen Elasticsearch-Index mit dem Alias `pokemon_index` zu erstellen 
```bash
go build -o command -tags=command && ./command build pokemon_index
```

## Testing
Im Browser die URL http://127.0.0.1:9221 öffnen.

Output
```json
{
  "name" : "lunch_n_learn",
  "cluster_name" : "lunch_n_learn_cluster",
  "cluster_uuid" : "0TyykRgaQE6Dfs4lOIozwg",
  "version" : {
    "number" : "8.7.0",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "09520b59b6bc1057340b55750186466ea715e30e",
    "build_date" : "2023-03-27T16:31:09.816451435Z",
    "build_snapshot" : false,
    "lucene_version" : "9.5.0",
    "minimum_wire_compatibility_version" : "7.17.0",
    "minimum_index_compatibility_version" : "7.0.0"
  },
  "tagline" : "You Know, for Search"
}
```

## Erste Schritte
Im Browser die URL http://127.0.0.1:5611/app/dev_tools#/console öffnen.
Damit öffnet sich Kibana mit den Dev-Tools.

Für Queries befinden sich in der Datei `Elasticsearch.md` unter dem Punkt Queries Beispiele für den Datensatz.