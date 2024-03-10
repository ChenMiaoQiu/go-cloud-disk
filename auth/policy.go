package auth

import "github.com/ChenMiaoQiu/go-cloud-disk/model"

func initPolicy() {
	// add base policies
	Casbin.AddPolicies(
		[][]string{
			// suspend user can't do anything
			{model.StatusSuspendUser, "*", "*", "deny"},
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
			{model.StatusAdmin, "admin/user*", "*", "allow"},
			{model.StatusAdmin, "admin/login*", "*", "allow"},
			{model.StatusAdmin, "admin/filestore*", "*", "allow"},
			{model.StatusAdmin, "admin/share*", "*", "allow"},
			{model.StatusAdmin, "admin/file*", "*", "allow"},
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
