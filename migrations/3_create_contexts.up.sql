CREATE TABLE "contexts" (
  "id" bigserial,
  "user_id" bigint NOT NULL REFERENCES users("id"),
  "text" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);

ALTER TABLE "actions"
  ADD COLUMN "context_id" bigint REFERENCES contexts("id");
