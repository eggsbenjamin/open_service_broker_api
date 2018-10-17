CREATE TABLE service (
  id          SERIAL PRIMARY KEY NOT NULL,
  service_id  UUID NOT NULL,
  name        VARCHAR(512) NOT NULL,
  description VARCHAR(1024) NOT NULL,
  tags        JSONB NOT NULL,
  requires    JSONB NOT NULL
);
