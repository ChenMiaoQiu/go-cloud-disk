package auth

import (
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Casbin *casbin.Enforcer

func InitCasbin() {
	a, err := gormadapter.NewAdapter("mysql", os.Getenv("MYSQL_DSN"), true)
	if err != nil {
		panic(err)
	}

	m, err := model.NewModelFromString(`
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act, eft
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow)) && !some(where (p.eft == deny))
	
	[matchers]
	m = g(r.sub, p.sub) && keyMatch(r.act, p.act) && keyMatch(r.obj, p.obj)
	`)
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}
	Casbin = e

	Casbin.LoadPolicy()

	if ok, _ := Casbin.Enforce("common_admin", "admin/user", "POST"); !ok {
		initPolicy()
	}
}
