-- Creation of product table
CREATE TYPE wallet_type AS ENUM ('Savings', 'Credit Card', 'Crypto Wallet');

CREATE TABLE IF NOT EXISTS user_wallet (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	user_name VARCHAR(255) NOT NULL,
	wallet_name VARCHAR(255) NOT NULL,
	wallet_type wallet_type NOT NULL,
	balance DECIMAL(10, 2) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user_wallet (user_id, user_name, wallet_name, wallet_type, balance) VALUES
(1, 'John Doe', 'John Savings', 'Savings', 1000.00),
(1, 'John Doe', 'John Credit Card', 'Credit Card', 500.00),
(1, 'John Doe', 'John Crypto Wallet', 'Crypto Wallet', 100.00),
(2, 'Jane Doe', 'Jane Savings', 'Savings', 2000.00),
(2, 'Jane Doe', 'Jane Credit Card', 'Credit Card', 1000.00),
(2, 'Jane Doe', 'Jane Crypto Wallet', 'Crypto Wallet', 200.00);

