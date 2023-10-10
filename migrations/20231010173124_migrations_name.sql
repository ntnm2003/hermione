CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE ROLE_TYPE as Enum('ADMIN','USER');
CREATE TABLE IF NOT EXISTS "users" (
    id uuid not null default uuid_generate_v4(),
    full_name varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    role ROLE_TYPE not null,
    created_at timestamp without time zone default now(),
    updated_at timestamp without time zone default now(),
    primary key (id)
    );
CREATE TABLE IF NOT EXISTs "vendors" (
    id uuid not null default uuid_generate_v4(),
    vendor varchar(255) not null,
    created_at timestamp without time zone default now(),
    updated_at timestamp without time zone default now(),
    primary key (id)
    );
CREATE TABLE IF NOT EXISTs "items" (

    id uuid not null default uuid_generate_v4(),
    item varchar(255) not null,
    price int not null,
    exp date not null,
    vendor_id uuid REFERENCES vendors(id),
    remaining_amount int,
    sold_amount int,
    created_at timestamp without time zone default now(),
    updated_at timestamp without time zone default now(),
    primary key (id)
    )
