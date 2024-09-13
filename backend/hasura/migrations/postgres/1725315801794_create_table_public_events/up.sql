CREATE TABLE "public"."events" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "title" text NOT NULL, "description" text NOT NULL, "images" jsonb NOT NULL, "featured_image" text NOT NULL, "venue" text NOT NULL, "address" text NOT NULL, "price" text NOT NULL, "specific_price" numeric NOT NULL, "preparation_date" date NOT NULL, "category" text NOT NULL, "tags" text[] NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "user_id" uuid NOT NULL DEFAULT gen_random_uuid(), PRIMARY KEY ("id") , FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE restrict ON DELETE restrict, CONSTRAINT "catagory_constraint" CHECK (category IN ('food', 'tech', 'education', 'entertainment', 'sport')), CONSTRAINT "price_constraint" CHECK (price IN('free','paid')));
CREATE OR REPLACE FUNCTION "public"."set_current_timestamp_updated_at"()
RETURNS TRIGGER AS $$
DECLARE
  _new record;
BEGIN
  _new := NEW;
  _new."updated_at" = NOW();
  RETURN _new;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER "set_public_events_updated_at"
BEFORE UPDATE ON "public"."events"
FOR EACH ROW
EXECUTE PROCEDURE "public"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_public_events_updated_at" ON "public"."events"
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
