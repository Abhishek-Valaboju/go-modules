CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    balance NUMERIC(10,2) DEFAULT 0.00,
    version INT DEFAULT 1
);