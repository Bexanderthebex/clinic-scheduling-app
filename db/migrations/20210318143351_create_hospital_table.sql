-- migrate:up
CREATE TABLE hospitals (
    id uuid PRIMARY KEY,
    name VARCHAR(100),
    city VARCHAR(100),
    address VARCHAR(100),
    lat NUMERIC(10, 8),
    long NUMERIC(11, 8),
    logo VARCHAR(300)
);

-- migrate:down
DROP TABLE hospitals;
