package query

const (
	// CREATE
	CreateCryptocurrency = `
		INSERT INTO cryptocurrencies (symbol, name, image_url, description, supply, max_supply)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	// READ
	GetCryptocurrencyByID = `
		SELECT id, symbol, name, image_url, description, supply, max_supply
		FROM cryptocurrencies
		WHERE id = $1
	`

	GetCryptocurrencyBySymbol = `
		SELECT id, symbol, name, image_url, description, supply, max_supply
		FROM cryptocurrencies
		WHERE symbol = $1
	`

	GetAllCryptocurrencies = `
		SELECT id, symbol, name, image_url, description, supply, max_supply
		FROM cryptocurrencies
	`

	// UPDATE
	UpdateCryptocurrency = `
		UPDATE cryptocurrencies
SET 
    name = $2, 
    description = $3, 
    supply = $4, 
    max_supply = $5, 
    image_url = $6, 
    updated_at = CURRENT_TIMESTAMP
WHERE symbol = $1
RETURNING id, symbol, name, description, supply, max_supply;

	`

	// DELETE
	DeleteCryptocurrency = `
		DELETE FROM cryptocurrencies
		WHERE symbol = $1
	`
)
