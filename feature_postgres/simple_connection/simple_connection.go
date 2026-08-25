package simpleconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	server_conn := os.Getenv("SERVER_CONN")
	return pgx.Connect(ctx, server_conn)
}
