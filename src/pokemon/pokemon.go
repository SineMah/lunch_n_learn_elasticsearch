package pokemon

type Entry struct {
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

type Ability struct {
	Name     string
	Slot     int
	IsHidden bool
}

type Pokemon struct {
	ID         int
	Name       string
	Games      []string
	Species    string
	Type       []string
	Abilities  []Ability
	Attributes Attributes
	Moves      []string
	Height     int
	Weight     int
	Images     []string
}
