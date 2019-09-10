package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"shop/config"
)

var (
	DbConn *sql.DB
	err error
)

//初始话
func init() {
	
	DbConn, err = sql.Open("mysql", config.Conn)
	if err != nil {
		panic(err.Error())
	}
}