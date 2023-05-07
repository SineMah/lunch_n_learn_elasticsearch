package routes

type Version struct {
	App     string `json:"app"`
	Version string `json:"version"`
}

type Pokemon struct {
	ID    int      `json:"ID"`
	Name  string   `json:"Name"`
	Games []string `json:"Games"`
	//Species    string       `json:"Species"`
	Type       []string     `json:"Type"`
	Abilities  []Ability    `json:"Abilities"`
	Attributes PokemonStats `json:"Attributes"`
	Moves      []string     `json:"Moves"`
	Height     int          `json:"Height"`
	Weight     int          `json:"Weight"`
	Images     []string     `json:"Images"`
}

type Ability struct {
	Name     string `json:"Name"`
	Slot     int    `json:"Slot"`
	IsHidden bool   `json:"IsHidden"`
}

type PokemonStats struct {
	HP             int `json:"HP"`
	Attack         int `json:"Attack"`
	Defense        int `json:"Defense"`
	SpecialAttack  int `json:"SpecialAttack"`
	SpecialDefense int `json:"SpecialDefense"`
	Speed          int `json:"Speed"`
}
