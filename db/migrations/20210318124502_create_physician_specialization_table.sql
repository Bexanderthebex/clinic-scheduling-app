-- migrate:up
CREATE TABLE PhysicianSpecialization (
    Id INT PRIMARY KEY,
    PhysicianId uuid,
    Specialization INT
);

-- migrate:down
DROP TABLE PhysicianSpecialization;
