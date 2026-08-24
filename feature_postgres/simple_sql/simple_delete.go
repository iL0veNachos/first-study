package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, idTask int) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = $1;
	`
	_, err := conn.Exec(ctx, sqlQuery, idTask)
	return err
}
