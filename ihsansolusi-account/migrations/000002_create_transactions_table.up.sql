CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT,
    sender_id INT,
    receiver_id INT,
    amount NUMERIC NOT NULL,
    type VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP
);