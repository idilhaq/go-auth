/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE TABLE users (
	id serial PRIMARY KEY,
  phone_number VARCHAR(50) UNIQUE NOT NULL,
  full_name VARCHAR(50),
  password VARCHAR,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE activity (
	id serial PRIMARY KEY,
  user_id int,
  last_login TIMESTAMP WITH TIME ZONE,
  login_attempt int,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE,
  CONSTRAINT unique_user_id UNIQUE (user_id)
);