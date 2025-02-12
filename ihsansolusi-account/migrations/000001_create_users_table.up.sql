CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nik VARCHAR(16) UNIQUE NOT NULL,
    phone VARCHAR(15) UNIQUE NOT NULL,
    account_number VARCHAR(15) UNIQUE NOT NULL,
    name VARCHAR(15) NOT NULL,
    balance NUMERIC DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP
);