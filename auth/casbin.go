package auth

import (
	"fmt"
	"os"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Casbin *casbin.Enforcer

func initPolicy() {
	// active user
	Casbin.AddPolicies(
		[][]string{
			// inactive user can't create file and filefolder
			{model.StatusInactiveUser, "user*", "*", "allow"},
			{model.StatusInactiveUser, "file*", "GET", "allow"},
			{model.StatusInactiveUser, "file*", "DELETE", "allow"},
			{model.StatusInactiveUser, "filefolder*", "GET", "allow"},
			{model.StatusInactiveUser, "filefolder*", "DELETE", "allow"},
			{model.StatusInactiveUser, "filestore*", "GET", "allow"},
			{model.StatusInactiveUser, "share*", "GET", "allow"},
			{model.StatusInactiveUser, "share*", "DELETE", "allow"},
			// active user can create file, filefolder and share
			{model.StatusActiveUser, "share*", "*", "allow"},
			{model.StatusActiveUser, "file*", "*", "allow"},
			{model.StatusActiveUser, "filefolder*", "*", "allow"},
			{model.StatusActiveUser, "rank*", "GET", "allow"},
			// admin user can change user status

			// super admin can do anything
			{model.StatusSuperAdmin, "*", "*", "allow"},
		},
	)

	// add group policies
	Casbin.AddGroupingPolicies(
		[][]string{
			{model.StatusActiveUser, model.StatusInactiveUser},
			{model.StatusAdmin, model.StatusActiveUser},
		},
	)
	Casbin.SavePolicy()
}

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

	if ok, err := Casbin.Enforce(model.StatusSuperAdmin, "share", "POST"); !ok {
		fmt.Println("create policy ", err)
		initPolicy()
	}
}
