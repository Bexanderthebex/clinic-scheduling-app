-- migrate:up
CREATE TABLE physician_specializations (
    physician_id uuid,
    specialization_id uuid
);

-- migrate:down
DROP TABLE physician_specializations;
