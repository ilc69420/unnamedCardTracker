package pokemon

import (
	"cardforge/internal/app"
	"context"
	"database/sql"
)

type PokemonService struct {
	App *app.App
}

func (s *PokemonService) getCard(ctx context.Context, card PokemonCard) (PokemonCard, error) {
	var fetchedCard PokemonCard

	stmt := `
	SELECT card_name, card_number, card_set
	FROM pokemon_cards
	WHERE card_name = ? AND card_number = ? AND card_set = ?
	LIMIT 1;
	`
	row := s.App.DB.QueryRowContext(
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

func (s *PokemonService) insertAuctions(ctx context.Context, auctions []PokemonAuction) error {
	tx, err := s.App.DB.Begin()
	if err != nil {
		s.App.Logger.DB("could not start transaction for insertAuctions function", "")
		return err
	}

	for _, card := range auctions {
		var cardID int
		err := s.App.DB.QueryRowContext(ctx, `
    SELECT card_id 
    FROM pokemon_cards 
    WHERE card_name = ? AND card_number = ? AND card_set = ?`,
			card.Name, card.Number, card.Set,
		).Scan(&cardID)

		if err == sql.ErrNoRows {
			res, err := tx.ExecContext(ctx, `
        INSERT OR IGNORE INTO pokemon_cards (card_name, card_number, card_set)
        VALUES (?, ?, ?)`,
				card.Name, card.Number, card.Set,
			)
			if err != nil {
				tx.Rollback()
				return err
			}
			insertedID, _ := res.LastInsertId()
			cardID = int(insertedID)
		} else if err != nil {
			tx.Rollback()
			return err
		}

		stmt := `
		INSERT OR IGNORE INTO pokemon_auctions
		(card_id, price, seller, status)
		VALUES
		(?,?,?,?)
		`
		if _, err = tx.ExecContext(ctx, stmt, cardID, card.Price, card.Seller, card.Status); err != nil {
			s.App.Logger.DB(err.Error(), stmt)
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (s *PokemonService) postCard(ctx context.Context, cards []PokemonCard) error {
	tx, err := s.App.DB.Begin()
	if err != nil {
		s.App.Logger.DB("could not start transaction on pokemon_cards", "")
		return err
	}
	stmt := `
	INSERT OR IGNORE INTO pokemon_cards 
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

			s.App.Logger.DB("couldn insert into pokemon_cards", stmt)
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
