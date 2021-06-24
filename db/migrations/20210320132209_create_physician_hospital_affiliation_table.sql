-- migrate:up
CREATE TABLE physician_hospitals(
    physician_id uuid,
    hospital_id uuid,
    start_date timestamp,
    end_date timestamp
);

-- migrate:down
DROP TABLE physician_hospitals;
