CREATE TABLE service_instance (
  id          SERIAL PRIMARY KEY NOT NULL,
  plan_id     INTEGER REFERENCES plan(id) NOT NULL,
  context     JSONB NOT NULL,
  parameters  JSONB NOT NULL
);
