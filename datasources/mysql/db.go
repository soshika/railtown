package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"railtown/logger"
)

var (
	Client *sql.DB
)

func init() {

	username := "root"
	password := "root"
	host := "localhost"
	schema := "railtown_DB"

	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = Client.Ping(); err != nil {
		fmt.Println(err.Error())
	}
	logger.Info("database successfully configured!")
}
