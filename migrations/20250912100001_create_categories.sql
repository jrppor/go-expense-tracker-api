-- +goose Up
CREATE TABLE IF NOT EXISTS categories (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NULL,
  name VARCHAR(64) NOT NULL,
  slug VARCHAR(80) NOT NULL,
  description VARCHAR(255),
  color VARCHAR(7),
  icon VARCHAR(64),
  sort_order INT DEFAULT 0,
  parent_id BIGINT NULL,
  budget_limit_cents BIGINT NOT NULL DEFAULT 0,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ NULL
);

-- unique ภายในขอบเขต user_id (หรือ global ถ้า user_id เป็น NULL เหมือนกัน)
CREATE UNIQUE INDEX IF NOT EXISTS ux_user_name ON categories (COALESCE(user_id, 0), name);
CREATE UNIQUE INDEX IF NOT EXISTS ux_user_slug ON categories (COALESCE(user_id, 0), slug);

CREATE INDEX IF NOT EXISTS idx_categories_user_id ON categories (user_id);
CREATE INDEX IF NOT EXISTS idx_categories_parent_id ON categories (parent_id);
CREATE INDEX IF NOT EXISTS idx_categories_deleted_at ON categories (deleted_at);

-- +goose Down
DROP TABLE IF EXISTS categories;