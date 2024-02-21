-- +goose Up
CREATE TABLE IF NOT EXISTS ideas (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    track TEXT NOT NULL,
    is_selected BOOLEAN NOT NULL DEFAULT false
);

-- +goose Down

DROP TABLE ideas;