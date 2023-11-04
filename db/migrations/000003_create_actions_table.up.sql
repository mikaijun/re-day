BEGIN;
  CREATE TABLE IF NOT EXISTS actions(
    id UUID NOT NULL PRIMARY KEY,
    task_id UUID,
    score INTEGER NOT NULL,
    comment TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY(task_id)
    REFERENCES tasks(id)  
  );
  CREATE INDEX on tasks(id);
COMMIT;
