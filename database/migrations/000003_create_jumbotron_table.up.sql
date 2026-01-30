CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE jumbotrons (
    internal_id BIGSERIAL PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text,
    image text,
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    public_id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT jumbotron_public_id_unique UNIQUE (public_id)
);