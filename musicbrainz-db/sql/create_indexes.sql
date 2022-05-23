\set ON_ERROR_STOP 1
BEGIN;

CREATE INDEX IF NOT EXISTS artist_idx_trgm ON artist USING gin(name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS artist_idx_trgm_sort ON artist USING gin(sort_name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS artist_alias_idx_trgm ON artist_alias USING gin(name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS artist_alias_idx_trgm_sort ON artist_alias USING gin(sort_name gin_trgm_ops);

CREATE INDEX IF NOT EXISTS label_idx_trgm ON label USING gin(name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS label_alias_idx_trgm ON label_alias USING gin(name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS label_alias_idx_trgm_sort ON label_alias USING gin(sort_name gin_trgm_ops);

COMMIT;
