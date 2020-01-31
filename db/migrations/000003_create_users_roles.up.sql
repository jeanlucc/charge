CREATE TABLE IF NOT EXISTS roles(
   id serial PRIMARY KEY,
   role VARCHAR (255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS roles_users_relation(
   role_id int REFERENCES roles (id) ON UPDATE CASCADE ON DELETE CASCADE,
   user_id int REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
   CONSTRAINT roles_users_relation_pkey PRIMARY KEY (role_id, user_id)
);
