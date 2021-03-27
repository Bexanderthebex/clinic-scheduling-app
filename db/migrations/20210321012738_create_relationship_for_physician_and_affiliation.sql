-- migrate:up
ALTER TABLE physician_specializations
    ADD CONSTRAINT fk_physician_specializations_physician_id
        FOREIGN KEY (physician_id)
        REFERENCES physicians (id),
    ADD CONSTRAINT fk_physician_specializations_specialization_id
        FOREIGN KEY (specialization_id)
        REFERENCES specializations (id);

-- migrate:down
ALTER TABLE physician_specializations
    DROP CONSTRAINT fk_physician_specializations_physician_id,
    DROP CONSTRAINT fk_physician_specializations_specialization_id;
