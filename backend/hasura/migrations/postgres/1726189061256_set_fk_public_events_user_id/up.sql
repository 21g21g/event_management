alter table "public"."events" drop constraint "events_user_id_fkey",
  add constraint "events_user_id_fkey"
  foreign key ("user_id")
  references "public"."users"
  ("id") on update cascade on delete cascade;
