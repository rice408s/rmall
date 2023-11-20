package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"rmall/config"
)

func main() {
	DBtest()
}

type test struct {
	Id int
	D1 string
}

func DBtest() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("init DB failed, err:", err)
		return
	}
	sqlStr := `select * from test where id = :id`
	var t test
	query, err := db.NamedQuery(sqlStr, map[string]interface{}{"id": "1"})
	if err != nil {
		return
	}
	for query.Next() {
		err := query.StructScan(&t)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)
	}
}
