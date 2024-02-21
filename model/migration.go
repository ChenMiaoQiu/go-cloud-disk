package model

// migration database
func migration() {
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&File{})
	_ = DB.AutoMigrate(&FileFolder{})
	_ = DB.AutoMigrate(&FileStore{})
}
