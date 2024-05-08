package postgres

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	c := mysql.Config{
		DBName:    "clean_architecture",
		User:      "user",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Loc:       jst,
	}
	db, err = sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return err
	}
	return nil
}
