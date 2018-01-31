CREATE TABLE users
(
  id         UUID PRIMARY KEY ,
  email      VARCHAR(255) NOT NULL UNIQUE,
  password   BYTEA NOT NULL,
  ip_address VARCHAR(45),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);