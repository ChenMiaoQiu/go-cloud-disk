package model

import (
	"os"
)

// migration database
func migration() {
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&File{})
	_ = DB.AutoMigrate(&FileFolder{})
	_ = DB.AutoMigrate(&FileStore{})
	_ = DB.AutoMigrate(&Share{})
	initSuperAdmin()
}

func initSuperAdmin() {
	// create super admin
	var count int64
	adminUserName := os.Getenv("ADMIN_USERNAME")
	if err := DB.Model(&User{}).Where("user_name = ?", adminUserName).Count(&count).Error; err != nil {
		panic("create super admin err %v")
	}

	if count == 0 {
		if err := createSuperAdmin(); err != nil {
			panic("create super admin err %v")
		}
	}
}
