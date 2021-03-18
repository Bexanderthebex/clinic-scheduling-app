-- migrate:up
CREATE TABLE Hospital (
    Id uuid PRIMARY KEY,
    Name VARCHAR(100),
    City VARCHAR(100),
    Address VARCHAR(100),
    Lat NUMERIC(10, 8),
    Long NUMERIC(10, 8),
    Logo VARCHAR(300)
);

-- migrate:down
DROP TABLE Hospital;
