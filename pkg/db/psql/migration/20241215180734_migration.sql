-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE users (
                       id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                       first_name VARCHAR(50) NOT NULL,
                       last_name VARCHAR(50) NOT NULL,
                       username VARCHAR(50) NOT NULL UNIQUE,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       password_hash VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username ON users(username);


CREATE INDEX idx_users_email ON users(email);


CREATE TABLE cryptocurrencies (
                                  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                  symbol VARCHAR(10) NOT NULL UNIQUE,
                                  image_url TEXT NOT NULL,
                                  name VARCHAR(100) NOT NULL,
                                  description TEXT,
                                  supply DECIMAL(38, 18),
                                  max_supply DECIMAL(38, 18),
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  updated_at timestamp DEFAULT current_timestamp
);

CREATE INDEX idx_cryptocurrencies_symbol ON cryptocurrencies(symbol);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cryptocurrencies CASCADE;

DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
