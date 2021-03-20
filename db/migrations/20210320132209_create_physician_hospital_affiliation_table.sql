-- migrate:up
CREATE TABLE PhysicianHospitalAffiliation(
    Id INT PRIMARY KEY,
    PhysicianId uuid,
    HospitalId uuid,
    StartDate timestamp,
    EndDate timestamp
);

-- migrate:down
DROP TABLE PhysicianHospitalAffiliation;
