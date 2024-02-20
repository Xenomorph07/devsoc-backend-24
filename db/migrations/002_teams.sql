-- +goose Up
CREATE TABLE IF NOT EXISTS teams (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    teamcode TEXT NOT NULL UNIQUE,
    memberid UUID[] NOT NULL,
    projectid UUID NOT NULL,
    ideaid UUID NOT NULL,
    round INTEGER NOT NULL DEFAULT 1
);

-- +goose Down

DROP TABLE teams;