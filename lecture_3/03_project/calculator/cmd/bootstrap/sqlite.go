package bootstrap

import (
	"database/sql"

	"code-cadets-2021/lecture_3/03_project/calculator/cmd/config"
	_ "github.com/mattn/go-sqlite3"
)

func Sqlite() *sql.DB {
	db, err := sql.Open("sqlite3", config.Cfg.SqliteDatabase)
	if err != nil {
		panic(err)
	}

	return db
}
