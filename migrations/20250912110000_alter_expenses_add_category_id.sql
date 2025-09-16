-- +goose Up
-- +goose StatementBegin
ALTER TABLE expenses
    ADD COLUMN IF NOT EXISTS category_id BIGINT NULL;

-- Backfill category_id by matching expenses.category to categories.name
UPDATE expenses e
SET category_id = c.id
FROM categories c
WHERE e.category_id IS NULL AND e.category = c.name;

-- Index and FK
CREATE INDEX IF NOT EXISTS idx_expenses_category_id ON expenses(category_id);
ALTER TABLE expenses
    ADD CONSTRAINT fk_expenses_category
    FOREIGN KEY (category_id) REFERENCES categories(id)
    ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE expenses DROP CONSTRAINT IF EXISTS fk_expenses_category;
DROP INDEX IF EXISTS idx_expenses_category_id;
ALTER TABLE expenses DROP COLUMN IF EXISTS category_id;
-- +goose StatementEnd
