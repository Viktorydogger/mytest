CREATE TABLE IF NOT EXISTS items(
id SERIAL PRIMARY KEY,
customer smallint NOT NULL,
services smallint NOT NULL,
orders smallint NOT NULL,
balance money NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);