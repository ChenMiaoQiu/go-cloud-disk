package auth

import (
	"fmt"
	"os"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Casbin *casbin.Enforcer

func InitCasbin() {
	a, err := gormadapter.NewAdapter("mysql", os.Getenv("MYSQL_DSN"), true)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("./auth/rbac_model.conf", a)
	if err != nil {
		panic(err)
	}
	Casbin = e

	Casbin.LoadPolicy()

	if ok, err := Casbin.Enforce(model.StatusAdmin, "admin/user", "POST"); !ok {
		fmt.Println("create policy ", err)
		initPolicy()
	}
}
