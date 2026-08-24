package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRow(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	taskSlice := []TaskModel{}
	sqlQuery := `
	SELECT title, description, completed, created_at, completed_at FROM tasks
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task TaskModel
		err := rows.Scan(
			&task.title,
			&task.description,
			&task.completed,
			&task.created_at,
			&task.completed_at,
		)
		taskSlice = append(taskSlice, task)
		if err != nil {
			return nil, err
		}
	}
	return taskSlice, nil
}
