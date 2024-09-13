alter table "public"."users" alter column "image_id" set default gen_random_uuid();
alter table "public"."users"
  add constraint "users_image_id_fkey"
  foreign key (image_id)
  references "public"."images"
  (id) on update restrict on delete restrict;
alter table "public"."users" alter column "image_id" drop not null;
alter table "public"."users" add column "image_id" uuid;
