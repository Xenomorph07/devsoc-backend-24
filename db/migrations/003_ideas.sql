-- +goose Up
CREATE TABLE IF NOT EXISTS ideas (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    track TEXT NOT NULL,
    github TEXT NOT NULL default '',
    figma TEXT NOT NULL default '',
    others TEXT NOT NULL default '',
    is_selected BOOLEAN NOT NULL DEFAULT false,
    teamid UUID NOT NULL UNIQUE
);

-- +goose Down

DROP TABLE ideas;