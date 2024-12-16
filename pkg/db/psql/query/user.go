package query

const (
	// Создать пользователя
	CreateUser = `
		INSERT INTO users (first_name, last_name, username, email, password_hash)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	// Получить пользователя по UUID
	GetUserByUUID = `
		SELECT id, first_name, last_name, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	// Получить пользователя по username
	GetUserByUsername = `
		SELECT id, first_name, last_name, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	// Получить пользователя по email
	GetUserByEmail = `
		SELECT id, first_name, last_name, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	// Обновить пользователя
	UpdateUser = `
		UPDATE users
		SET first_name = $2, last_name = $3, username = $4, email = $5, password_hash = $6, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`

	// Удалить пользователя
	DeleteUser = `
		DELETE FROM users
		WHERE id = $1
	`
)
