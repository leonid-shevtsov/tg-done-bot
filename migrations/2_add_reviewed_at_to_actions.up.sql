ALTER TABLE actions ADD "reviewed_at" timestamptz;
UPDATE actions SET "reviewed_at" = "created_at";
ALTER TABLE actions ALTER "reviewed_at" SET NOT NULL;
