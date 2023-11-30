package initialize

import (
	"fmt"
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"rmall/global"
)

func initCasbin() {
	//初始化casbin
	a, err := sqlxadapter.NewAdapter(global.Db, "casbin_rule")
	if err != nil {
		panic(err)
	}
	//加载策略
	global.E, err = casbin.NewEnforcer("casbin_model.conf", a)
	if err != nil {
		panic(err)
	}
	user, err := global.E.AddRoleForUser("yyy", "admin")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	forUser, err := global.E.GetRolesForUser("xxx")
	if err != nil {
		panic(err)
	}
	fmt.Println(forUser)
	role, err := global.E.GetUsersForRole("admin")
	if err != nil {
		panic(err)
	}

	fmt.Println(role)
	//fmt.Println(service.GetPolicy())
	//fmt.Println(service.GetPolicyByRole("hw"))
	//fmt.Println(service.GetPolicyByPath("/api/v1/casbin"))
	//fmt.Println(service.GetPolicyByMethod("POST"))
	//if err = e.LoadPolicy(); err != nil {
	//	log.Println("casbin.LoadPolicy failed,err:", err)
	//}
	//
	////检测权限
	//has, err := e.Enforce("alice", "data1", "read")
	//if err != nil {
	//	log.Println("Enforce failed, err: ", err)
	//}
	//fmt.Println(has)
	//
	////policy, err := e.AddPolicy("hw", "/api/v1/casbin", "POST")
	////if err != nil {
	////	log.Println("AddPolicy failed, err: ", err)
	////}
	////fmt.Println(policy)
	//fmt.Println(e.GetPolicy())
	//
	//has, err = e.Enforce("admin", "/api/v1/casbin", "POST")
	//if err != nil {
	//	log.Println("Enforce failed, err: ", err)
	//}
	//fmt.Println(has)
	//
	//// Modify the policy.
	//// e.AddPolicy(...)
	//// e.RemovePolicy(...)
	//
	//// Save the policy back to DB.
	//if err = e.SavePolicy(); err != nil {
	//	log.Println("SavePolicy failed, err: ", err)
	//}
	//return nil
}
