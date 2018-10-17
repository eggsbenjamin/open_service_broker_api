CREATE TABLE service_plan (
  id          SERIAL PRIMARY KEY NOT NULL,
  name        VARCHAR(512) NOT NULL,
  service_id  INTEGER REFERENCES service(id)
);
