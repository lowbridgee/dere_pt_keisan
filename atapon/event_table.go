package atapon

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

type EventTable struct {
	Id int64 `db:"pk" column:"event_id"`
	Name string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Point int64
	Item int64
}

func CreateDB() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, ":memory:")
    // db, err := genmai.New(&genmai.MySQLDialect{}, "dsn")
    // db, err := genmai.New(&genmai.PostgresDialect{}, "dsn")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    if err := db.CreateTable(&EventTable{}); err != nil {
        panic(err)
    }
}

func InsertTest() {
	obj := &EventTable{
		Name: "atapon",Point: 20000, Item: 3000
	}
}
