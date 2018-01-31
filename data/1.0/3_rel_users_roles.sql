CREATE TABLE rel_users_roles (
  user_id UUID REFERENCES users (id) ON UPDATE CASCADE,
  role_id UUID REFERENCES roles (id) ON UPDATE CASCADE,
  CONSTRAINT users_roles_pkey PRIMARY KEY (user_id, role_id)
);
