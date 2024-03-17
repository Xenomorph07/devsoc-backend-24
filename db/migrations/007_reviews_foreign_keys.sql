-- +goose Up
ALTER TABLE reviews
ADD CONSTRAINT fk_reviews
FOREIGN KEY (team_id)
REFERENCES teams(id)
ON UPDATE CASCADE;

-- +goose Down
ALTER TABLE reviews
DROP CONSTRAINT fk_reviews;
