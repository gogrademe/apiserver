
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE userAccount (
  id SERIAL PRIMARY KEY,
  email text NOT NULL,
  emailLower text NOT NULL,
  role text,
  hashedPassword text,
  createdAt TIMESTAMP NOT NULL,
  updatedAt TIMESTAMP NOT NULL,
  disabled boolean,
  UNIQUE (email, emailLower)
);

-- +goose Down
DROP TABLE userAccount;
-- SQL section 'Down' is executed when this migration is rolled back

