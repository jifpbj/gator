-- +goose Up

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT,
    url TEXT NOT NULL UNIQUE,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID
);

-- +goose Down
DROP TABLE posts;

-- postgres://jif:@localhost:5432/gator
