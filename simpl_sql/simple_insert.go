package simplsql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn pgx.Conn) error {
	sqlQwery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ('gome', 'сделать', 'FALSE', '2025-11-26 18:32:02')
	`
	_, err := conn.Exec(ctx, sqlQwery)
	return err
}
