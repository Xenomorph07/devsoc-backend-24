-- +goose Up
CREATE TABLE IF NOT EXISTS projects(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    github TEXT NOT NULL default '',
    figma TEXT NOT NULL default '',
    track TEXT NOT NULL default '',
    teamid UUID NOT NULL UNIQUE
);

-- +goose Down

DROP TABLE projects;