-- migrate:up
ALTER TABLE qualifications
    ADD CONSTRAINT fk_qualifications_physician_id
        FOREIGN KEY (physician_id)
            REFERENCES physicians (Id)
            ON DELETE CASCADE
            ON UPDATE CASCADE;

-- migrate:down
ALTER TABLE qualifications
    DROP CONSTRAINT fk_qualifications_physician_id;