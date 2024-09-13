alter table "public"."bookmarks" drop constraint "bookmarks_event_id_fkey",
  add constraint "bookmarks_event_id_fkey"
  foreign key ("event_id")
  references "public"."events"
  ("id") on update restrict on delete restrict;
