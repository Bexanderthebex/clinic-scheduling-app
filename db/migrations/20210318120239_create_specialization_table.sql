-- migrate:up
CREATE TABLE Specialization (
    Id INT PRIMARY KEY,
    UniqueCode uuid,
    SpecializationName varchar(100)
);

-- migrate:down
DROP TABLE Specialization;
