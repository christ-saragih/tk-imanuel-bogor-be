ALTER TABLE blogs DROP CONSTRAINT IF EXISTS blogs_slug_key;
ALTER TABLE blogs DROP CONSTRAINT IF EXISTS blogs_title_key;

DROP INDEX IF EXISTS blogs_slug_key;
DROP INDEX IF EXISTS blogs_title_key;

-- Unique hanya jika deleted_at IS NULL
CREATE UNIQUE INDEX blogs_slug_unique_active ON blogs (slug) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX blogs_title_unique_active ON blogs (title) WHERE deleted_at IS NULL;
