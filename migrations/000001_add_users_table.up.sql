-- +migrate Up
CREATE TABLE IF NOT EXISTS user_info (
    id SERIAL PRIMARY KEY,
    email VARCHAR(30) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(10) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE
);