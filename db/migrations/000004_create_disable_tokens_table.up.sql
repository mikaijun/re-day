BEGIN;
  CREATE TABLE IF NOT EXISTS disable_tokens(
    token text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
  );
 CREATE INDEX on disable_tokens(token);
COMMIT;
