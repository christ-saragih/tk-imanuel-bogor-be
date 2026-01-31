CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE galleries (
    internal_id BIGSERIAL PRIMARY KEY,
    
    title VARCHAR(255) NOT NULL,
    image TEXT,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    public_id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT gallery_public_id_unique UNIQUE (public_id)
);