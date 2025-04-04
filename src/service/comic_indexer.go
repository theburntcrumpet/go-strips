package service

import (
	"os"

	"time"

	"github.com/google/uuid"
	"github.com/theburntcrumpet/go-strip/src/config"
	"github.com/theburntcrumpet/go-strip/src/database"
	"github.com/theburntcrumpet/go-strip/src/model"
	"gorm.io/gorm"
)

type ComicIndexer interface {
	IndexComicsAtPath(path string) error
	indexComic(path string) error
	IndexComics() error
}

type ComicIndexerImpl struct {
	db        gorm.DB
	conf      config.ServiceConfig
	cbzParser ComicParser
}

func NewComicIndexer(dbFactory database.DatabaseFactory, conf config.ServiceConfig) (ComicIndexer, error) {
	db, err := dbFactory.CreateDatabase()
	if err != nil {
		return nil, err
	}
	return &ComicIndexerImpl{db: *db, conf: conf, cbzParser: &CbzComicParser{}}, nil
}

func (ci *ComicIndexerImpl) indexComic(path string) error {
	comic := &model.Comic{}
	ci.db.FirstOrInit(comic, "filename = ?", path)
	if comic.Filename != "" {
		return nil
	}
	os.MkdirAll(ci.conf.AppDataDirectory+"/previews", os.ModePerm)

	previewKey, err := ci.writePreviewImage(path)
	if err != nil {
		return err
	}
	comic = &model.Comic{
		Filename:        path,
		Progress:        0,
		PreviewImageKey: previewKey,
		LastOpenedTime:  time.Now().Format(time.RFC3339),
		CreatedTime:     time.Now().Format(time.RFC3339),
	}
	ci.db.Create(comic)
	return nil
}

func (ci *ComicIndexerImpl) IndexComics() error {
	return ci.IndexComicsAtPath(ci.conf.ComicDirectory)
}

func (ci *ComicIndexerImpl) IndexComicsAtPath(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			err = ci.IndexComicsAtPath(path + "/" + entry.Name())
			if err != nil {
				return err
			}
		} else {
			err = ci.indexComic(path + "/" + entry.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ci *ComicIndexerImpl) writePreviewImage(path string) (string, error) {
	parser := ci.cbzParser
	comic, err := parser.ParseComic(path)
	if err != nil {
		return "", err
	}
	if len(comic.Pages) == 0 {
		return "", nil
	}
	uuidStr := uuid.NewString()
	previewPath := ci.conf.AppDataDirectory + "/previews/" + uuidStr + ".jpg"
	err = os.WriteFile(previewPath, comic.Pages[0].Data, 0644)
	if err != nil {
		return "", err
	}
	return uuidStr, nil
}
