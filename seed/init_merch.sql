INSERT INTO merch (id, name, price)
VALUES
    (gen_random_uuid(), 't-shirt', 80),
    (gen_random_uuid(), 'cup', 20),
    (gen_random_uuid(), 'book', 50),
    (gen_random_uuid(), 'pen', 10),
    (gen_random_uuid(), 'powerbank', 200),
    (gen_random_uuid(), 'hoody', 300),
    (gen_random_uuid(), 'umbrella', 200),
    (gen_random_uuid(), 'socks', 10),
    (gen_random_uuid(), 'wallet', 50),
    (gen_random_uuid(), 'pink-hoody', 500)
ON CONFLICT (id) DO NOTHING;
