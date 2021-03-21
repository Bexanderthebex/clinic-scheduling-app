-- migrate:up
ALTER TABLE PhysicianHospitalAffiliation
    ADD CONSTRAINT FK_PhysicianHospitalAffiliation_PhysicianId
        FOREIGN KEY (PhysicianId)
            REFERENCES Physician (id),
    ADD CONSTRAINT FK_PhysicianHospitalAffiliation_HospitalId
        FOREIGN KEY (HospitalId)
            REFERENCES Hospital (id);

-- migrate:down
ALTER TABLE PhysicianHospitalAffiliation
    DROP CONSTRAINT FK_PhysicianHospitalAffiliation_PhysicianId,
    DROP CONSTRAINT FK_PhysicianHospitalAffiliation_HospitalId;
