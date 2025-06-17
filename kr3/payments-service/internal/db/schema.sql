DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    user_id TEXT UNIQUE NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


DROP TABLE IF EXISTS inbox;

CREATE TABLE inbox (
    id SERIAL PRIMARY KEY,
    message_id UUID NOT NULL,
    topic TEXT NOT NULL,
    payload JSONB NOT NULL,
    processed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now()
);