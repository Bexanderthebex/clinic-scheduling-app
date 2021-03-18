-- migrate:up
CREATE TABLE Qualification (
    Id INT PRIMARY KEY,
    PhysicianId uuid,
    QualificationName VARCHAR(100),
    InstituteName VARCHAR(100),
    ProcurementYear TIMESTAMP
);

-- migrate:down
DROP TABLE Qualification;
