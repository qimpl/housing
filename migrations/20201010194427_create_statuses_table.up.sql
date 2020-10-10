CREATE TABLE "statuses" (
  "id" uuid PRIMARY KEY UNIQUE DEFAULT gen_random_uuid (),
  "name" varchar(50) NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_timestamp
  BEFORE UPDATE ON "statuses"
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_update_timestamp ();

ALTER TABLE "housings"
  ADD COLUMN "status_id" uuid REFERENCES "statuses" (id) NOT NULL;
