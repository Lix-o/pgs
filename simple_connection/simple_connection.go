package simpleconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	ConString := os.Getenv("CONNSTRING")
	return pgx.Connect(ctx, ConString)
}
