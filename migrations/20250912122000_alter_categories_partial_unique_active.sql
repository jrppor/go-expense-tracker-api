-- +goose Up
-- +goose StatementBegin
-- Replace global unique indexes with partial unique indexes that only apply to active (non-deleted) rows
DROP INDEX IF EXISTS ux_categories_name;
DROP INDEX IF EXISTS ux_categories_slug;

CREATE UNIQUE INDEX IF NOT EXISTS ux_categories_name ON categories (name) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS ux_categories_slug ON categories (slug) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Revert back to global unique indexes
DROP INDEX IF EXISTS ux_categories_name;
DROP INDEX IF EXISTS ux_categories_slug;

CREATE UNIQUE INDEX IF NOT EXISTS ux_categories_name ON categories (name);
CREATE UNIQUE INDEX IF NOT EXISTS ux_categories_slug ON categories (slug);
-- +goose StatementEnd
