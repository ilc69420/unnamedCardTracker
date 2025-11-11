package pokemon

import (
	"cardforge/internal/app"
	"encoding/json"
	"net/http"
	"net/url"
)

type Pokemon struct {
	App     *app.App
	Service *PokemonService
}

func (m *Pokemon) RegisterPokemonRoutes(app *app.App) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /card", m.getCard)
	r.HandleFunc("POST /card", m.postCard)

	return r
}

func (h *Pokemon) getCard(w http.ResponseWriter, r *http.Request) {
	// Extract queryparams -> fetch from db -> send card as json
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	cardName := queryParams.Get("name")
	cardNum := queryParams.Get("number")
	cardSet := queryParams.Get("set")

	card := PokemonCard{
		Name:   cardName,
		Number: cardNum,
		Set:    cardSet,
	}

	cardFromDb, err := h.Service.getCard(r.Context(), card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cardFromDb); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Pokemon) postCard(w http.ResponseWriter, r *http.Request) {
	// extract card from req body -> insert into db -> return 201
	var cards []PokemonCard
	if err := json.NewDecoder(r.Body).Decode(&cards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Service.postCard(r.Context(), cards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
