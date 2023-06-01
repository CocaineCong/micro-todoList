package model

func migration() {
	// 自动迁移模式
	_db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&Task{})
}
