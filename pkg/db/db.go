package db

import "github.com/go-pg/pg/v10"

func ConnectDB() *pg.DB {
	opt, err := pg.ParseURL("postgres://postgres:postgres@localhost:15434/go_fiber?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	return db
}
