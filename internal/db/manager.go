package db

import (
	"fmt"
	"github.com/PabloGilvan/transaction/internal/config/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseManager interface {
	GetDatabaseConnection() (*gorm.DB, error)
	Migrate(model interface{}) error
}

type DatabaseManagerImpl struct {
}

func NewDatabaseManager() DatabaseManager {
	return DatabaseManagerImpl{}
}

func (manager DatabaseManagerImpl) GetDatabaseConnection() (*gorm.DB, error) {
	var host = global.Viper.GetString("app.database.host")
	var user = global.Viper.GetString("app.database.user")
	var pass = global.Viper.GetString("app.database.password")
	var database = global.Viper.GetString("app.database.name")
	var port = global.Viper.GetString("app.database.port")
	var sslmode = global.Viper.GetString("app.database.sslmode")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, pass, database, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func (manager DatabaseManagerImpl) Migrate(model interface{}) error {
	db, err := manager.GetDatabaseConnection()
	if err != nil {
		return err
	}

	return db.AutoMigrate(model)
}
