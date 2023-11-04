BEGIN;
  CREATE TABLE IF NOT EXISTS tasks(
    id VARCHAR (255) UNIQUE NOT NULL PRIMARY KEY,
    content TEXT,
    user_id VARCHAR (255),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
  );
  CREATE INDEX on tasks(id);
COMMIT;
