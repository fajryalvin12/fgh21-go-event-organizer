create table "profile" (
    "id" serial primary key,
    "picture" varchar(255),
    "full_name" varchar(80),
    "birth_date" varchar(20),
    "gender" smallint,
    "phone_number" varchar(15),
    "profession" varchar(80),
    "nationality_id" int references "nationalities"("id"),
    "user_id" int references "users"("id") 
);