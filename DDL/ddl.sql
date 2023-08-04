-- Tabel customers
CREATE TABLE customers (
    id INT NOT NULL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel merchants
CREATE TABLE merchants (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    account_number VARCHAR(100) NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel transactions
CREATE TABLE transactions (
    id INT PRIMARY KEY,
    customer_id INT,
    merchant_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    transaction_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);

