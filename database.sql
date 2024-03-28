CREATE SCHEMA IF NOT EXISTS core;

CREATE TABLE core.users (
  id serial PRIMARY KEY,
  name varchar(60),
  phone_number varchar(13),
  password_hash bytea,
  password_salt bytea
);

CREATE INDEX ix_users_phone_number ON core.users USING btree (phone_number);
CREATE INDEX ix_users_name ON core.users USING btree (name);
