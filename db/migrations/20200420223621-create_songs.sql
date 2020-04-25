
-- +migrate Up
CREATE TABLE IF NOT EXISTS songs (
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
--     id int PRIMARY KEY AUTO_INCREMENT NOT NULL,
--     name text NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255),
--     date timestamp NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS songs;