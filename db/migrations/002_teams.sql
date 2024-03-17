-- +goose Up
CREATE TABLE IF NOT EXISTS teams (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    code TEXT NOT NULL UNIQUE,
    leader_id UUID NOT NULL,
    --members_id UUID[] NOT NULL,
    round INTEGER NOT NULL DEFAULT 1
);

-- +goose Down
DROP TABLE IF EXISTS teams;