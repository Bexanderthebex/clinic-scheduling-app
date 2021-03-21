-- migrate:up
ALTER TABLE PhysicianSpecialization
    ADD CONSTRAINT FK_PhysicianSpecialization_PhysicianId
        FOREIGN KEY (PhysicianId)
        REFERENCES Physician (id),
    ADD CONSTRAINT FK_PhysicianSpecialization_SpecializationId
        FOREIGN KEY (SpecializationId)
        REFERENCES Specialization (id);

-- migrate:down
ALTER TABLE PhysicianSpecialization
    DROP CONSTRAINT FK_PhysicianSpecialization_PhysicianId,
    DROP CONSTRAINT FK_PhysicianSpecialization_SpecializationId;
