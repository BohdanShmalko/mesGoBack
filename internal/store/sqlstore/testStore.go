package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			truncateString := "TRUNCATE " + strings.Join(tables, ", ") + " CASCADE"
			if _, err := db.Exec(truncateString); err != nil {
				t.Fatal(err)
			}
		}
		db.Close()
	}
}
