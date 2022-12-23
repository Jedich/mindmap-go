package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mindmap-go/utils/config"
	"strings"
)

type Database struct {
	Connection *gorm.DB
	Config     *config.Config
	Log        *zap.Logger
}

func NewDatabase(cfg *config.Config, log *zap.Logger) *Database {
	db := &Database{
		Config: cfg,
		Log:    log,
	}

	return db
}

func (db *Database) OpenConnection() {
	var err error
	switch s := strings.ToLower(db.Config.DB.Driver); s {
	case "mysql":
		db.Connection, err = gorm.Open(mysql.Open(db.Config.DB.MySQL.DSN), &gorm.Config{
			//Logger: logger.New(
			//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			//	logger.Config{
			//		SlowThreshold: time.Second,
			//		LogLevel:      logger.Info,
			//		Colorful:      true,
			//	}),
		})
		if err != nil {
			db.Log.Fatal(fmt.Sprintf("%+v", err.Error()))
		}
	default:
		db.Log.Error(fmt.Sprintf("Unsupported driver %s", s))
	}
}

func (db *Database) CloseConnection() {
	sqlDB, err := db.Connection.DB()
	if err != nil {
		db.Log.Error(err.Error())
	}
	sqlDB.Close()
}
