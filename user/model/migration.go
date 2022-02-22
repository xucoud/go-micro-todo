package model

func migration() {
	Db.Set(`gorm:table_options`, "charset=utf8").
		AutoMigrate(&User{})
}
