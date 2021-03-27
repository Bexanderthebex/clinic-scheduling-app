-- migrate:up transaction:false
CREATE TABLE IF NOT EXISTS physicians (
    id uuid PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR,
    middle_name VARCHAR
);

-- migrate:down transaction:false
DROP TABLE physicians;
