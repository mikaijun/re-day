BEGIN;
  CREATE TABLE IF NOT EXISTS tasks(
    id UUID NOT NULL PRIMARY KEY,
    content TEXT,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY(user_id)
    REFERENCES users(id) 
  );
  CREATE INDEX on tasks(id);
COMMIT;
