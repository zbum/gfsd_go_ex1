package common

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DataSource struct {
	db *sql.DB
}

func NewDataSource() *DataSource {
	d := &DataSource{}
	d.initialize()
	return d
}

func (d *DataSource) initialize() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "test",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "gfsd",
		AllowNativePasswords: true,
	}
	var err error
	d.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = d.db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DataSource) GetDB() *sql.DB {
	return d.db
}
