package pokemon

type PokemonCard struct {
	Name   string `json:"card_name"`
	Number string `json:"card_number"`
	Set    string `json:"card_set"`
}

type PokemonAuction struct {
	PokemonCard
	Seller string `json:"seller"`
	Status bool   `json:"status"`
	Price  int    `json:"price"`
}
