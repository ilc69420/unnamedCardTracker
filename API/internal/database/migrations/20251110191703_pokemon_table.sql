-- +goose Up
-- +goose StatementBegin
CREATE TABLE pokemon_cards (
  card_id INTEGER PRIMARY KEY AUTOINCREMENT,
  card_name TEXT NOT NULL,
  card_number TEXT NOT NULL,
  card_set TEXT NOT NULL,
  UNIQUE(card_name, card_number, card_set)
);

CREATE TABLE pokemon_auctions (
  auction_id INTEGER PRIMARY KEY AUTOINCREMENT,
  card_id INTEGER NOT NULL,
  price INTEGER NOT NULL,
  seller TEXT NOT NULL,
  status BOOLEAN NOT NULL,
  FOREIGN KEY (card_id) REFERENCES pokemon_cards(card_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pokemon_cards;
DROP TABLE IF EXISTS pokemon_auctions;
-- +goose StatementEnd
