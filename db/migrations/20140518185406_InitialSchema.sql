
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

CREATE TABLE class (
  id SERIAL PRIMARY KEY,
  name text NOT NULL,
  grade_level text NOT NULL,
  subject text NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE TABLE class_term (
  id SERIAL PRIMARY KEY,
  class_id integer PRIMARY KEY REFERENCES class,
  name text NOT NULL,
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE class_term_people (
  class_term_id integer PRIMARY KEY REFERENCES class,
  person_id integer PRIMARY KEY REFERENCES person,
  role text NOT NULL
  enrolled_date TIMESTAMP NOT NULL,
  unenrolled_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);


CREATE TABLE person (
  id SERIAL PRIMARY KEY,
  first_name text NOT NULL,
  middle_name text,
  last_name text NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE assignment (
  id SERIAL PRIMARY KEY,
  class_term_id integer PRIMARY KEY REFERENCES class_term,
  name text NOT NULL,
  type text NOT NULL,
  due_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE student_profile (
  person_id integer PRIMARY KEY REFERENCES person,
  grade_level text NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE TABLE teacher_profile (
  person_id integer PRIMARY KEY REFERENCES person,
  phone_number text NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);


-- +goose Down
DROP TABLE user_account;
DROP TABLE class;
DROP TABLE class_term;
DROP TABLE class_term_people;
DROP TABLE person;
DROP TABLE student_profile;
DROP TABLE teacher_profile;
-- SQL section 'Down' is executed when this migration is rolled back
