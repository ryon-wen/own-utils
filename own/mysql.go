package own

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQL struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func InitModel(sql *MySQL) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local&parseTime=True&charset=utf8mb4", sql.User, sql.Password, sql.Host, sql.Port, sql.Database)), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("connect mysql failed: " + err.Error())
	}

	// Test connection
	err = db.Exec("SELECT 1").Error
	if err != nil {
		panic("Failed to ping database:" + err.Error())
	}
	zap.S().Debug("Successfully connected to MySQL database!")
	return db
}

// Todo: auto rollback or commit
//func AutoTx(txF func(tx *gorm.DB) error) error {
//	tx := db.Begin()
//	err := txF(tx)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//	return nil
//}
