CREATE EXTENSION IF NOT EXISTS pgcrypto;
alter table "public"."imagestore" add column "event_id" uuid
 null default gen_random_uuid();
