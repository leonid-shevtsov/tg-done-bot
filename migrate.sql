CREATE TABLE "inbox_items" ("id" bigserial, "user_id" bigint NOT NULL, "text" text NOT NULL, "created_at" timestamptz NOT NULL, "processed_at" timestamptz, PRIMARY KEY ("id"))
CREATE TABLE "users" ("id" bigserial, "state" bigint NOT NULL, "current_inbox_item_id" bigint, "current_goal_id" bigint, "created_at" timestamptz NOT NULL, "last_message_at" timestamptz NOT NULL, PRIMARY KEY ("id"))
CREATE TABLE "goals" ("id" bigserial, "user_id" bigint NOT NULL, "text" text NOT NULL, "created_at" timestamptz NOT NULL, "completed_at" timestamptz, "dropped_at" timestamptz, PRIMARY KEY ("id"))
CREATE TABLE "actions" ("id" bigserial, "user_id" bigint NOT NULL, "goal_id" bigint NOT NULL, "text" text NOT NULL, "created_at" timestamptz NOT NULL, "completed_at" timestamptz, PRIMARY KEY ("id"))
