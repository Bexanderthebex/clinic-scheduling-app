-- migrate:up
CREATE TABLE qualifications (
    id INT PRIMARY KEY,
    physician_id uuid,
    qualification_name VARCHAR(100),
    institute_name VARCHAR(100),
    procurement_year TIMESTAMP
);

-- migrate:down
DROP TABLE qualifications;
