
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
--     name varchar(255) NOT NULL,
    name varchar(255),
    email varchar(255) NOT NULL,
    age int,
    password varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS users;