CREATE TABLE rel_users_roles
(
  user_id INTEGER REFERENCES users,
  role_id INTEGER REFERENCES roles
);

