package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gog/conf"
	"log"
)

func main() {
	var db *sql.DB
	db, err := sql.Open(conf.DBDialect, conf.DBConnStr)
	if err != nil {
		log.Output(0, "gog:connet to database fail")
		log.Fatalln(err)
	}
	defer db.Close()
}
