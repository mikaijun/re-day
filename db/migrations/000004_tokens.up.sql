BEGIN;
  CREATE TABLE IF NOT EXISTS tokens(
    id UUID NOT NULL PRIMARY KEY,
    token VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY(user_id)
    REFERENCES users(id) 
  );
 CREATE INDEX on tokens(id);
COMMIT;
