# Go Expense Tracker API

A small Go + Gin + GORM API for tracking expenses with PostgreSQL, Docker Compose, and Goose migrations.

## Prerequisites
- Docker and Docker Compose
- Make sure port 8080 is free on your machine (or set `PORT` env)

## Quick Start
```bash
# Start database and API
docker compose up -d db
docker compose up -d api

# Check container status
docker compose ps

# Smoke test
curl -sS http://localhost:8080/expenses
```

You can also use the HTTP request file in VS Code: open `http/expenses.http` and click "Send Request" on each block.

## Configuration
- Environment comes from Docker Compose (DB host `db`, default port 5432, DB `expensedb`).
- App port defaults to 8080; override with `PORT` env.
- `config/config.yml` can set `server.port` and is merged with env vars.

## Migrations (Goose)
Migrations are written in SQL and applied automatically on API startup.

- Place files in `migrations/` with timestamp naming, e.g.:
  - `migrations/20250912100000_create_expenses.sql`
  - `migrations/20250912100001_create_categories.sql`
  - `migrations/20250912110000_alter_expenses_add_category_id.sql`
- Restart API to apply new migrations:
Note on categories uniqueness:
- The API enforces unique `name` and `slug` only among active categories (soft deletes do not block). This is implemented via partial unique indexes (`WHERE deleted_at IS NULL`).
- If you POST a duplicate `name` or `slug` that already exists and is active, the API returns HTTP 409 Conflict.

```bash
docker compose restart api
docker compose logs -f api
```

## Common Commands
```bash
# Follow logs
docker compose logs -f db
docker compose logs -f api

# Rebuild/restart API after code changes
docker compose build api
docker compose up -d api

# List tables
docker compose exec db psql -U admin -d expensedb -c "\\dt*"

# Inspect data
docker compose exec -e PGPASSWORD=password db psql -U admin -d expensedb -c "SELECT * FROM expenses LIMIT 10;"
docker compose exec -e PGPASSWORD=password db psql -U admin -d expensedb -c "SELECT * FROM categories LIMIT 10;"

# Verify category linkage
docker compose exec -e PGPASSWORD=password db psql -U admin -d expensedb -c \
"SELECT e.id, e.amount, e.category_id, c.name AS category_name, e.date
   FROM expenses e
   LEFT JOIN categories c ON c.id = e.category_id
   ORDER BY e.date DESC, e.id DESC
   LIMIT 10;"

# Seed sample categories (safe to re-run)
docker compose exec -e PGPASSWORD=password db psql -U admin -d expensedb -c \
"INSERT INTO categories (name, slug, created_at, updated_at) VALUES
 ('อาหาร','food', NOW(), NOW()),
 ('เดินทาง','travel', NOW(), NOW())
 ON CONFLICT DO NOTHING;"

# Insert sample expense linked to category_id = 1
docker compose exec -e PGPASSWORD=password db psql -U admin -d expensedb -c \
"INSERT INTO expenses (category_id, amount, description, payment_method, tags, date, created_at, updated_at)
 VALUES (1, 150.75, 'ข้าวมันไก่', 'เงินสด', 'lunch', '2025-09-05', NOW(), NOW());"

# Demonstrate soft delete (uses deleted_at)
docker compose exec -e PGPASSWORD=password db psql -U admin -d expensedb -c \
"UPDATE expenses SET deleted_at = NOW() WHERE id = 1;"
docker compose exec -e PGPASSWORD=password db psql -U admin -d expensedb -c \
"SELECT id, deleted_at FROM expenses WHERE id = 1;"
```

## HTTP Requests
- Use `http/expenses.http` for quick tests (GET/POST/PUT/DELETE).
- API base URL: `http://localhost:8080`

## Troubleshooting
- API 404 or cannot connect: ensure the API runs at `http://localhost:8080`.
- DB not ready: check health and logs `docker compose logs -f db`.
- Schema not updated: verify Goose logs in `docker compose logs -f api` and try `docker compose restart api`.
- Foreign key errors with `category_id`: ensure a matching `categories.id` exists.

## Project Layout
- `cmd/` entrypoint
- `config/` configuration and DB init
- `controllers/`, `services/`, `repositories/` layers
- `models/` GORM models
- `migrations/` SQL migrations
- `routes/` HTTP routes (Gin)
- `http/` HTTP client request samples

## License
MIT (or update as appropriate)