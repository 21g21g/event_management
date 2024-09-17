alter table "public"."tickets" alter column "purchase_date" drop not null;
alter table "public"."tickets" add column "purchase_date" timestamptz;
