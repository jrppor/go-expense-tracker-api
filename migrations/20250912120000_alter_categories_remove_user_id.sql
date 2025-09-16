-- +goose Up
-- +goose StatementBegin
-- Drop user-scoped unique indexes first
DROP INDEX IF EXISTS ux_user_name;
DROP INDEX IF EXISTS ux_user_slug;
DROP INDEX IF EXISTS idx_categories_user_id;

-- Drop the user_id column
ALTER TABLE categories DROP COLUMN IF EXISTS user_id;

-- Optionally add global unique constraints for name and slug
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_indexes WHERE schemaname = current_schema() AND indexname = 'ux_categories_name'
    ) THEN
        CREATE UNIQUE INDEX ux_categories_name ON categories (name);
    END IF;
    IF NOT EXISTS (
        SELECT 1 FROM pg_indexes WHERE schemaname = current_schema() AND indexname = 'ux_categories_slug'
    ) THEN
        CREATE UNIQUE INDEX ux_categories_slug ON categories (slug);
    END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Remove global unique indexes added in Up
DROP INDEX IF EXISTS ux_categories_name;
DROP INDEX IF EXISTS ux_categories_slug;

-- Re-add user_id column
ALTER TABLE categories ADD COLUMN IF NOT EXISTS user_id BIGINT NULL;

-- Recreate original indexes
CREATE UNIQUE INDEX IF NOT EXISTS ux_user_name ON categories (COALESCE(user_id, 0), name);
CREATE UNIQUE INDEX IF NOT EXISTS ux_user_slug ON categories (COALESCE(user_id, 0), slug);
CREATE INDEX IF NOT EXISTS idx_categories_user_id ON categories (user_id);
-- +goose StatementEnd
