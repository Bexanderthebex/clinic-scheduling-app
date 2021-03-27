-- migrate:up
CREATE TABLE specializations (
    id INT PRIMARY KEY,
    unique_code uuid,
    specialization_name varchar(100)
);

-- migrate:down
DROP TABLE specializations;
