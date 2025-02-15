DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
            CREATE TYPE status AS ENUM ('success', 'failure');
        END IF;
    END $$;

CREATE TABLE IF NOT EXISTS users
(
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    balance  INTEGER      NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions
(
    id                 UUID                     DEFAULT gen_random_uuid() PRIMARY KEY,
    from_user          VARCHAR(255) REFERENCES users (username) ON DELETE CASCADE,
    to_user            VARCHAR(255) REFERENCES users (username) ON DELETE CASCADE,
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
    username   VARCHAR(255) REFERENCES users (username) ON DELETE CASCADE,
    merch_name VARCHAR(255) REFERENCES merch (name) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);



