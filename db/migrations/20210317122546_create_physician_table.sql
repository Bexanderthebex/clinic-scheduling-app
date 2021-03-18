-- migrate:up transaction:false
CREATE TABLE IF NOT EXISTS Physician (
    Id uuid PRIMARY KEY,
    FirstName VARCHAR,
    LastName VARCHAR,
    MiddleName VARCHAR
);

-- migrate:down transaction:false
DROP TABLE Physician;
