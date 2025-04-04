package main

import (
	"fmt"

	"github.com/theburntcrumpet/go-strip/src/config"
	"github.com/theburntcrumpet/go-strip/src/database"
	"github.com/theburntcrumpet/go-strip/src/service"
)

func main() {
	conf, err := config.LoadServiceConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbFactory := database.NewDatabaseFactory(*conf)
	indexer, err := service.NewComicIndexer(dbFactory, *conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = indexer.IndexComics()
	if err != nil {
		fmt.Println(err)
		return
	}
}
