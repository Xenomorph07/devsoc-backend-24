-- +goose Up
CREATE TABLE IF NOT EXISTS reviews (
    id UUID PRIMARY KEY,
    team_id UUID NOT NULL,
    reviewer TEXT NOT NULL,
    innovation_score DOUBLE PRECISION NOT NULL,
    functionality_score DOUBLE PRECISION NOT NULL,
    design_score DOUBLE PRECISION NOT NULL,
    tech_score DOUBLE PRECISION NOT NULL,
    presentation_score DOUBLE PRECISION NOT NULL,
    comments TEXT NOT NULL,
    total_score DOUBLE PRECISION NOT NULL,
    review_round INTEGER NOT NULL  
);

-- +goose Down
DROP TABLE IF EXISTS reviews;
