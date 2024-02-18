package model

// migration database
func migration() {
	_ = DB.AutoMigrate(&User{})
}
