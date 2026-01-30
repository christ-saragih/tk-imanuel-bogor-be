DROP INDEX IF EXISTS blogs_slug_unique_active;
DROP INDEX IF EXISTS blogs_title_unique_active;

ALTER TABLE blogs ADD CONSTRAINT blogs_slug_key UNIQUE (slug);
ALTER TABLE blogs ADD CONSTRAINT blogs_title_key UNIQUE (title);
