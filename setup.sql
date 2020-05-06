DROP TABLE IF EXISTS records;

CREATE TABLE records (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id INT NOT NULL,
  subject_id INT,
  study_time INT NOT NULL,
  date_time TIMESTAMP NOT NULL
);
