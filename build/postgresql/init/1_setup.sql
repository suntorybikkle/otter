DROP TABLE IF EXISTS study_infos;

CREATE TABLE study_infos (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id INT NOT NULL,
  subject_id INT,
  study_time INT NOT NULL,
  date_time TIMESTAMP NOT NULL
);
