-- migrate:up
ALTER TABLE Qualification
    ADD CONSTRAINT FK_Qualification_PhysicianId
        FOREIGN KEY (PhysicianId)
            REFERENCES Physician (Id)
            ON DELETE CASCADE
            ON UPDATE CASCADE;

-- migrate:down
ALTER TABLE Qualification
    DROP CONSTRAINT FK_Qualification_PhysicianId;