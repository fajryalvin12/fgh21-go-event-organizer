create table "event_categories" (
    "id" serial primary key,
    "event_id" int references "events"("id"),
    "categories_id" int references "categories"("id")
);