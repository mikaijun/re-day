BEGIN;
  CREATE TABLE IF NOT EXISTS auth_expiries(
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY(user_id)
    REFERENCES users(id)
  );
 CREATE INDEX on auth_expiries(id);
COMMIT;
