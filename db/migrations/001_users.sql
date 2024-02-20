-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    reg_no TEXT NOT NULL,
    password TEXT NOT NULL,
    phone TEXT NOT NULL,
    college TEXT NOT NULL,
    gender TEXT NOT NULL,
    role TEXT NOT NULL,
    country TEXT NOT NULL,
    github TEXT NOT NULL,
    bio TEXT NOT NULL,
    is_banned BOOLEAN NOT NULL,
    is_added BOOLEAN NOT NULL,
    is_vitian BOOLEAN NOT NULL,
    is_verified BOOLEAN NOT NULL,
    team_id UUID
);

-- +goose Down
DROP TABLE users;