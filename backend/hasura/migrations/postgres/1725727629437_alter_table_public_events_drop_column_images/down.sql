alter table "public"."events" alter column "images" drop not null;
alter table "public"."events" add column "images" _text;
