CREATE EXTENSION pgcrypto;

CREATE FUNCTION trigger_update_timestamp ()
  RETURNS TRIGGER
  AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TABLE "housing_types" (
  "id" uuid PRIMARY KEY UNIQUE DEFAULT gen_random_uuid (),
  "name" varchar(50) NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_timestamp
  BEFORE UPDATE ON "housing_types"
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_update_timestamp ();

CREATE TABLE "housings" (
  "id" uuid PRIMARY KEY UNIQUE DEFAULT gen_random_uuid (),
  "type_id" uuid REFERENCES "housing_types" (id) NOT NULL,
  "surface_area" float NOT NULL,
  "rent_price" float NOT NULL,
  "rental_charges" float NOT NULL,
  "country" varchar(2) NOT NULL,
  "state" varchar(50) NOT NULL,
  "city" varchar(100) NOT NULL,
  "zip" varchar(12) NOT NULL,
  "street" varchar(255) NOT NULL,
  "is_furnished" bool DEFAULT FALSE NOT NULL,
  "has_electricity" bool DEFAULT FALSE NOT NULL,
  "has_gas" bool DEFAULT FALSE NOT NULL,
  "is_published" bool DEFAULT FALSE NOT NULL,
  "owner_id" uuid NOT NULL,
  "last_tenant_id" uuid,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_timestamp
  BEFORE UPDATE ON "housings"
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_update_timestamp ();
