alter table "public"."imagestore"
  add constraint "imagestore_event_id_fkey"
  foreign key ("event_id")
  references "public"."events"
  ("id") on update restrict on delete restrict;
