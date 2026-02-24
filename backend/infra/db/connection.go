package db

import (
	"fmt"
	"todolist/config"

	"github.com/jmoiron/sqlx"
)

func GetConnectionDB(cfg config.Config) string {
	return "user=fla password=FL@pon676701234 host=localhost port=5432 dbname=todo_db "
}

func NewConnection(cfg config.Config) (*sqlx.DB, error) {
	dbsource := GetConnectionDB(cfg)
	dbCon, err := sqlx.Connect("postgres", dbsource)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
