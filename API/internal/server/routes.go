package server

import (
	"cardforge/internal/app"
	"cardforge/internal/server/middleware"
	"cardforge/internal/server/pokemon"
	"net/http"
)

func registerRoutes(app *app.App) *http.ServeMux {
	r := http.NewServeMux()

	// Pokemon related
	pokemonHandler := &pokemon.Pokemon{
		App:     app,
		Service: &pokemon.PokemonService{App: app},
	}
	pokemonRoutes := pokemonHandler.RegisterPokemonRoutes(app)

	loggedRoutes := middleware.Logging(app.Logger, pokemonRoutes)

	r.Handle("/api/v1/", http.StripPrefix("/api/v1/pokemon", loggedRoutes))

	return r
}
