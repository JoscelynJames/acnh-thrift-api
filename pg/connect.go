package pg

import "github.com/go-pg/pg/v10"

func ConnectDB() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "user",
		Password: "pass",
		Database: "db_name",
	})
}
