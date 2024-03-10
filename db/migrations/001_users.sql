-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    reg_no TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    phone TEXT NOT NULL,
    college TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    country TEXT NOT NULL,
    gender TEXT NOT NULL,
    role TEXT NOT NULL,
    is_banned BOOLEAN NOT NULL,
    is_added BOOLEAN NOT NULL,
    is_vitian BOOLEAN NOT NULL,
    is_verified BOOLEAN NOT NULL,
    is_leader BOOLEAN NOT NULL,
    is_profile_complete BOOLEAN NOT NULL,
    team_id UUID
);

CREATE TABLE vit_details (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    email TEXT UNIQUE NOT NULL,
    block TEXT NOT NULL,
    room TEXT NOT NULL
);

-- +goose Down
DROP TABLE vit_details;
DROP TABLE users;
