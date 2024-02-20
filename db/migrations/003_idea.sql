-- +goose Up
CREATE TABLE IF NOT EXISTS ideas (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    video TEXT NOT NULL DEFAULT '',
    github TEXT NOT NULL DEFAULT '',
    figma TEXT NOT NULL DEFAULT ''
);

-- +goose Down

DROP TABLE ideas;