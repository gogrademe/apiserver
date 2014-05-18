
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_account (
  id SERIAL PRIMARY KEY,
  email text NOT NULL UNIQUE,
  email_lower text NOT NULL UNIQUE,
  role text,
  hashed_password text,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  disabled boolean DEFAULT FALSE
);

CREATE TABLE student_profile (
  id SERIAL PRIMARY KEY,
  grade_level text NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE TABLE person (
  id SERIAL PRIMARY KEY,
  first_name text NOT NULL,
  middle_name text,
  last_name text NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  student_profile_id integer REFERENCES student_profile
);


-- +goose Down
DROP TABLE user_account;
DROP TABLE person;
-- SQL section 'Down' is executed when this migration is rolled back

