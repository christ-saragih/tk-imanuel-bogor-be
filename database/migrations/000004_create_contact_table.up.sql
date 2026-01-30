CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE contacts (
    internal_id BIGSERIAL PRIMARY KEY,
    address TEXT,
    maps_embed_url TEXT,
    maps_link TEXT,
    email VARCHAR(255),
    phone_number VARCHAR(50),
    whatsapp_number VARCHAR(50),
    instagram_url VARCHAR(255),
    facebook_url VARCHAR(255),
    youtube_url VARCHAR(255),
    tiktok_url VARCHAR(255),
    opening_hours VARCHAR(255),
    last_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    public_id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT contact_public_id_unique UNIQUE (public_id)
);