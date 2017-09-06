CREATE TABLE "waiting_fors" (
  "id" bigserial,
  "user_id" bigint NOT NULL references users("id"),
  "goal_id" bigint NOT NULL references goals("id"),
  "text" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  "reviewed_at" timestamptz NOT NULL,
  "completed_at" timestamptz,
  "dropped_at" timestamptz,
  PRIMARY KEY ("id")
);

ALTER TABLE "users"
  ADD "current_waiting_for_id" bigint references waiting_fors("id");
