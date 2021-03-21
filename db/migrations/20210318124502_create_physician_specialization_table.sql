-- migrate:up
CREATE TABLE PhysicianSpecialization (
    Id INT PRIMARY KEY,
    PhysicianId uuid,
    SpecializationId INT
);

-- migrate:down
DROP TABLE PhysicianSpecialization;
