CREATE TABLE service_plan (
  id          SERIAL PRIMARY KEY NOT NULL,
  plan_id     UUID NOT NULL UNIQUE,
  name        VARCHAR(512) NOT NULL,
  service_id  INTEGER REFERENCES service(id)
);
