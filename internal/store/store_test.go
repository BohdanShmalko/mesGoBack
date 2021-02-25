package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(t *testing.M)  {

	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=mesuser password=1111 dbname=mesmusic_test sslmode=disable"
	}

	os.Exit(t.Run())
}
