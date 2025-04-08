package service

import (
	"os"

	"github.com/theburntcrumpet/go-strip/src/config"
	"github.com/theburntcrumpet/go-strip/src/database"
	"github.com/theburntcrumpet/go-strip/src/model"
	"gorm.io/gorm"
)

type ComicService interface {
	GetComics() ([]model.Comic, error)
	GetComicsPaginated(query string, page, pageSize int) ([]model.Comic, error)
	GetComicWithId(id string) (model.Comic, error)
	GetPreviewImageWithUuid(uuid string) ([]byte, error)
}

type ComicServiceImpl struct {
	db        gorm.DB
	conf      config.ServiceConfig
	cbzParser ComicParser
}

func NewComicService(dbFactory database.DatabaseFactory, conf config.ServiceConfig) (ComicService, error) {
	db, err := dbFactory.CreateDatabase()
	if err != nil {
		return nil, err
	}
	return &ComicServiceImpl{db: *db, conf: conf, cbzParser: &CbzComicParser{}}, nil
}

func (cs *ComicServiceImpl) GetComics() ([]model.Comic, error) {
	comics := []model.Comic{}
	err := cs.db.Find(&comics).Error
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func (cs *ComicServiceImpl) GetComicsPaginated(query string, page, pageSize int) ([]model.Comic, error) {
	comics := []model.Comic{}
	offset := (page - 1) * pageSize
	err := cs.db.Where("filename LIKE ?", "%"+query+"%").Offset(offset).Limit(pageSize).Find(&comics).Error
	if err != nil {
		return nil, err
	}
	return comics, nil
}

func (cs *ComicServiceImpl) GetComicWithId(id string) (model.Comic, error) {
	comic := model.Comic{}
	err := cs.db.Where("id = ?", id).First(&comic).Error
	if err != nil {
		return model.Comic{}, err
	}
	return comic, nil
}

func (cs *ComicServiceImpl) GetPreviewImageWithUuid(uuid string) ([]byte, error) {
	previewPath := cs.conf.AppDataDirectory + "/previews/" + uuid + ".jpg"
	previewImage, err := os.ReadFile(previewPath)
	if err != nil {
		return nil, err
	}
	return previewImage, nil
}
