package database

import (
	"github.com/theburntcrumpet/go-strip/src/config"
	"github.com/theburntcrumpet/go-strip/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseFactory interface {
	CreateDatabase() (*gorm.DB, error)
}

type DatabaseFactoryImpl struct {
	serviceConfig config.ServiceConfig
	hasMigrated   bool
}

func NewDatabaseFactory(config config.ServiceConfig) DatabaseFactory {
	return &DatabaseFactoryImpl{
		serviceConfig: config,
		hasMigrated:   false,
	}
}

func (d *DatabaseFactoryImpl) CreateDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(d.serviceConfig.AppDataDirectory+"/comic.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if !d.hasMigrated {
		err = db.AutoMigrate(&model.Comic{})
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
