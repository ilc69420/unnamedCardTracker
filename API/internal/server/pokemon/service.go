package pokemon

import (
	"context"
	"database/sql"
)

type PokemonService struct {
	DB *sql.DB
}

func (s *PokemonService) getCard(ctx context.Context, card PokemonCard) (PokemonCard, error) {
	var fetchedCard PokemonCard

	stmt := `
	SELECT card_name, card_number, card_set
	FROM pokemon_cards
	WHERE card_name = ? AND card_number = ? AND card_set = ?
	LIMIT 1;
	`
	row := s.DB.QueryRowContext(
		ctx,
		stmt,
		card.Name,
		card.Number,
		card.Set,
	)

	if err := row.Scan(
		&fetchedCard.Name,
		&fetchedCard.Number,
		&fetchedCard.Set,
	); err != nil {
		return PokemonCard{}, err
	}

	return fetchedCard, nil
}

func (s *PokemonService) postCard(ctx context.Context, cards []PokemonCard) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt := `
	INSERT INTO pokemon_cards 
	(card_name, card_number, card_set)
	VALUES
	(?,?,?)`

	for _, c := range cards {
		_, err := tx.ExecContext(
			ctx,
			stmt,
			c.Name,
			c.Number,
			c.Set,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
