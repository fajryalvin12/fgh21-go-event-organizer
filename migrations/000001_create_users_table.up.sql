create table "users" (
    "id" serial primary key,
    "email" varchar(80) unique not null,
    "password" varchar(100) not null,
    "username" varchar(50) unique
);