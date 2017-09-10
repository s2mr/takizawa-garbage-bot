-- +migrate Up
CREATE TABLE users (
  id SERIAL NOT NULL PRIMARY KEY ,
  user_id VARCHAR(256) NOT NULL,
  region INT,
  created timestamp NOT NULL DEFAULT NOW()
);
-- +migrate Down
DROP TABLE users;