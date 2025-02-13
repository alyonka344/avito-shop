CREATE TYPE status AS ENUM ('success', 'failure');

CREATE TABLE IF NOT EXISTS users
(
    id       UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    balance  INTEGER      NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions
(
    id                 UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    from_user_id       UUID REFERENCES users (id) ON DELETE CASCADE,
    to_user_id         UUID REFERENCES users (id) ON DELETE CASCADE,
    amount             INTEGER,
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    transaction_status status
);

CREATE TABLE IF NOT EXISTS merch
(
    id    UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name  VARCHAR(255) UNIQUE NOT NULL,
    price INTEGER             NOT NULL
);

CREATE TABLE IF NOT EXISTS purchases
(
    id         UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id    UUID REFERENCES users (id) ON DELETE CASCADE,
    merch_name VARCHAR(255) REFERENCES merch (name) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);



