package porsist


import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

var(
	Db *sql.DB
	err error
)

const (
	MaxConns int =100
	MixConns int = 2
)

func init()  {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/mydb?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	Db.SetMaxIdleConns(MaxConns)
	Db.SetMaxOpenConns(MixConns)
	err = Db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}