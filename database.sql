/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
create table "user"
(
	id uuid not null constraint user_pk primary key,
	phone_number varchar(13) not null constraint user_phone_number_unique unique,
	password varchar(150) not null,
	full_name varchar(60) not null,
	login_count integer default 0 not null,
	created_at timestamp default now() not null,
	updated_at timestamp
);

create index user_id_index on "user" (id);
create index user_phone_number_index on "user" (phone_number);
