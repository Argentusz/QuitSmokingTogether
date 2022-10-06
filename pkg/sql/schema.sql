CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    login TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    last_cig INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    users_list INTEGER[] DEFAULT 0,
    interval_time INTEGER NOT NULL DEFAULT 0
);

