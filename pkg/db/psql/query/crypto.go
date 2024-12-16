package query

const (
	// CREATE
	CreateCryptocurrency = `
		INSERT INTO cryptocurrencies (symbol, name, description, supply, max_supply)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	// READ
	GetCryptocurrencyByID = `
		SELECT id, symbol, name, description, supply, max_supply
		FROM cryptocurrencies
		WHERE id = $1
	`

	GetCryptocurrencyBySymbol = `
		SELECT id, symbol, name, description, supply, max_supply
		FROM cryptocurrencies
		WHERE symbol = $1
	`

	GetAllCryptocurrencies = `
		SELECT id, symbol, name, description, supply, max_supply
		FROM cryptocurrencies
	`

	// UPDATE
	UpdateCryptocurrency = `
		UPDATE cryptocurrencies
		SET symbol = $2, name = $3, description = $4, supply = $5, max_supply = $6, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
		RETURNING id, symbol, name, description, supply, max_supply
	`

	// DELETE
	DeleteCryptocurrency = `
		DELETE FROM cryptocurrencies
		WHERE symbol = $1
	`
)
