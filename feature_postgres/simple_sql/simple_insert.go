package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ('Domashka', 'To do math homework a', FALSE, '2025-11-26 18:01:05');
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
