package model

import (
	"log"
	"os"
	"time"

	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Database init mysql connect
func Database(connString string) {
	// init gorm log set
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)
	// connect database
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})

	if connString == "" || err != nil {
		loglog.Log().Error("mysql lost: %v", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		loglog.Log().Error("mysql lost: %v", err)
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	DB = db

	migration()
}
