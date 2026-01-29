CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE teachers (
    internal_id BIGSERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    slug varchar(100) UNIQUE,
    role varchar(255) NOT NULL DEFAULT 'teacher',
    photo text,
    color varchar(50),
    bio text,
    education text,
    experience SMALLINT DEFAULT 0,
    fun_fact text,
    quote text,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    public_id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT teacher_public_id_unique UNIQUE (public_id)
);