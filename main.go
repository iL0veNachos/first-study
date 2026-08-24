package main

import (
	"context"
	"fmt"
	simpleconnection "study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	if err := simple_sql.InsertRow(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("succeed")
}
