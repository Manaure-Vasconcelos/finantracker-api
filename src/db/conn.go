package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func ConnectDB() (*pgxpool.Pool, error) {

	db_url := "postgresql://finantracker:finantracker@localhost:5432/finantracker_db"

	dbpool, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to db")

	return dbpool, nil
}
