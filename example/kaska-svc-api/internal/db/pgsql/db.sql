
CREATE DATABASE kaska_db;
USE kaska_db;
CREATE TABLE IF NOT EXISTS kaska_users (
    not_user_id SERIAL,
    user_uuid VARCHAR(255) PRIMARY KEY NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR ( 255 ) UNIQUE,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
);

