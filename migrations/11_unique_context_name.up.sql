ALTER TABLE contexts ADD dropped_at timestamptz;

CREATE INDEX contexts_unique_text ON contexts (user_id, text) WHERE dropped_at IS NULL;
