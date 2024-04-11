package tutorial

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {

	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(pool)

	os.Exit(m.Run())
}
