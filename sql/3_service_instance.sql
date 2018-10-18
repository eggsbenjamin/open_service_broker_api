CREATE TABLE service_instance (
  id          SERIAL PRIMARY KEY NOT NULL,
  plan_id     INTEGER REFERENCES service_plan(id) NOT NULL,
  context     JSONB NOT NULL,
  parameters  JSONB NOT NULL
);
