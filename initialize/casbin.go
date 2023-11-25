package initialize

import (
	"fmt"
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"log"
	"rmall/global"
)

func InitCasbin() {
	a, err := sqlxadapter.NewAdapter(global.Db, "casbin_rule")
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("casbin_model.conf", a)
	if err != nil {
		panic(err)
	}
	if err = e.LoadPolicy(); err != nil {
		log.Println("casbin.LoadPolicy failed,err:", err)
	}

	has, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		log.Println("Enforce failed, err: ", err)
	}
	fmt.Println(has)

	//policy, err := e.AddPolicy("hw", "/api/v1/casbin", "POST")
	//if err != nil {
	//	log.Println("AddPolicy failed, err: ", err)
	//}
	//fmt.Println(policy)
	fmt.Println(e.GetPolicy())

	has, err = e.Enforce("admin", "/api/v1/casbin", "POST")
	if err != nil {
		log.Println("Enforce failed, err: ", err)
	}
	fmt.Println(has)

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	if err = e.SavePolicy(); err != nil {
		log.Println("SavePolicy failed, err: ", err)
	}
}
