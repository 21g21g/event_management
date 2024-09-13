alter table "public"."users"
  add constraint "users_image_id_fkey"
  foreign key ("image_id")
  references "public"."images"
  ("id") on update restrict on delete restrict;
