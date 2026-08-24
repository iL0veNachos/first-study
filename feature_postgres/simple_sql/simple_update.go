package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateCompleted(ctx context.Context, conn *pgx.Conn, id int, value bool) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = $1
	WHERE id = $2
	`
	_, err := conn.Exec(ctx, sqlQuery, value, id)
	return err
}
