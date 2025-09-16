-- +goose Up
-- +goose StatementBegin
-- Ensure category_id column and FK exist (from previous migration), then drop legacy 'category' column
ALTER TABLE expenses DROP COLUMN IF EXISTS category;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Re-add legacy column; attempt to backfill category name from categories by category_id
ALTER TABLE expenses ADD COLUMN IF NOT EXISTS category VARCHAR(100);

UPDATE expenses e
SET category = c.name
FROM categories c
WHERE e.category_id = c.id AND (e.category IS NULL OR e.category = '');
-- +goose StatementEnd