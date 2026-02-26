package db

import (
	"fmt"
	"todolist/config"

	"github.com/jmoiron/sqlx"
)

func GetConnectionDB(cfg *config.Config) string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
}

func NewConnection(cfg *config.Config) (*sqlx.DB, error) {
	dbsource := GetConnectionDB(cfg)
	dbCon, err := sqlx.Connect("postgres", dbsource)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
