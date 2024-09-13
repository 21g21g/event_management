CREATE EXTENSION IF NOT EXISTS pgcrypto;
alter table "public"."users" add column "image_id" uuid
 null default gen_random_uuid();
