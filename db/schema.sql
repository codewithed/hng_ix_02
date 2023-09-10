CREATE TABLE "persons" (
  "id" integer PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
