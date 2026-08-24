package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1, $2, $3, $4);
	`
	_, err := conn.Exec(
		ctx,
		sqlQuery,
		task.title,
		task.description,
		task.completed,
		task.created_at,
	)
	return err
}
