-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) UNIQUE NOT NULL,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE feeds;