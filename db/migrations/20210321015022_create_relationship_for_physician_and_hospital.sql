-- migrate:up
ALTER TABLE physician_hospitals
    ADD CONSTRAINT fk_physician_hospitals_physician_id
        FOREIGN KEY (physician_id)
            REFERENCES physicians (id),
    ADD CONSTRAINT fk_physician_hospitals_hospital_id
        FOREIGN KEY (hospital_id)
            REFERENCES hospitals (id);

-- migrate:down
ALTER TABLE physician_hospitals
    DROP CONSTRAINT fk_physician_hospitals_physician_id,
    DROP CONSTRAINT fk_physician_hospitals_hospital_id;
