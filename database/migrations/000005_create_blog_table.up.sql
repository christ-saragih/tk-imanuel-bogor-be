CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE blogs (
    internal_id BIGSERIAL PRIMARY KEY,
    
    slug VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(255) UNIQUE NOT NULL,
    excerpt TEXT,
    content TEXT,
    image TEXT,
    tags TEXT[], 
    view_count INTEGER DEFAULT 0,
    
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    public_id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT blog_public_id_unique UNIQUE (public_id)
);