-- migrate:up
CREATE TABLE specializations (
    id uuid PRIMARY KEY,
    specialization_name varchar(100),
    CONSTRAINT UK_specializations_specialization_name UNIQUE (specialization_name)
);

-- migrate:down
DROP TABLE specializations;
