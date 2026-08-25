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

	_, err = simple_sql.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}

	// for _, val := range rows {
	// 	val.Print()
	// }

	fmt.Println("succeed")
}
