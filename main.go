package main

import (
	simplsql "PostgreSQL/simpl_sql"
	simpleconnection "PostgreSQL/simple_connection"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simplsql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("все *хуенно")
}
