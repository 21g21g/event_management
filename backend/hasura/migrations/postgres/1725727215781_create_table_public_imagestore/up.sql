CREATE TABLE "public"."imagestore" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "url" text NOT NULL, PRIMARY KEY ("id") );
CREATE EXTENSION IF NOT EXISTS pgcrypto;
