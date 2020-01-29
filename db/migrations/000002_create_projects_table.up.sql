CREATE TABLE IF NOT EXISTS projects(
   id serial PRIMARY KEY,
   name VARCHAR (255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS projects_users_relation(
   project_id int REFERENCES projects (id) ON UPDATE CASCADE ON DELETE CASCADE,
   user_id int REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
   CONSTRAINT projects_users_relation_pkey PRIMARY KEY (project_id, user_id)
);
