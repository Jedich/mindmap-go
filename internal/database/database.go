package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mindmap-go/app/models"
	"mindmap-go/utils/config"
	"os"
	"strings"
	"time"
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
		retries := 3
		//var logger gorm
		l := logger.Default
		if db.Config.App.Production {
			l = logger.Default.LogMode(logger.Silent)
		}

		env, ok := os.LookupEnv("APP_DSN")
		if !ok {
			db.Log.Fatal("APP_DSN not set")
		}
		db.Log.Info(env)

		db.Connection, err = gorm.Open(mysql.Open(env), &gorm.Config{
			Logger: l,
		})
		for err != nil {
			db.Log.Info(fmt.Sprintf("Failed to connect to database (%d)", retries))
			if retries > 1 {
				retries--
				time.Sleep(5 * time.Second)
				db.Connection, err = gorm.Open(mysql.Open(env), &gorm.Config{
					Logger: l,
				})
				continue
			}
			db.Log.Fatal(err.Error())
		}
	default:
		db.Log.Fatal(fmt.Sprintf("Unsupported driver %s", s))
	}
	db.Log.Info("Connected to database")
	err = db.Connection.AutoMigrate(&models.Account{}, &models.User{}, &models.Card{}, &models.File{}, &models.Map{})
	if err != nil {
		db.Log.Fatal(err.Error())
	}
	db.Log.Info("Database migrated successfully")
}

func (db *Database) CloseConnection() {
	sqlDB, err := db.Connection.DB()
	if err != nil {
		db.Log.Fatal(err.Error())
	}
	sqlDB.Close()
}
