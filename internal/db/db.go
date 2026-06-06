package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/niviten/wfe/internal/config"
)

var dbconn *sql.DB

func Init() {
	userName := config.GetConfig("DB_USER")
	password := config.GetConfig("DB_PASS")
	host := config.GetConfig("DB_HOST")
	port := config.GetConfig("DB_PORT")
	dbname := config.GetConfig("DB_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, host, port, dbname)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Failed to open db: ", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalln("Failed to ping db: ", err)
	}

	dbconn = conn
}

func GetConn() *sql.DB {
	return dbconn
}
