alter table "public"."imagestore" drop constraint "imagestore_event_id_fkey",
  add constraint "imagestore_event_id_fkey"
  foreign key ("event_id")
  references "public"."events"
  ("id") on update restrict on delete restrict;
