CREATE TABLE "visits" (
  "id" uuid PRIMARY KEY UNIQUE DEFAULT gen_random_uuid (),
  "date" date NOT NULL,
  "hour" time NOT NULL,
  "is_accepted" bool DEFAULT FALSE NOT NULL,
  "housing_id" uuid REFERENCES "housings" (id) ON DELETE CASCADE NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_timestamp
  BEFORE UPDATE ON "visits"
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_update_timestamp ();
