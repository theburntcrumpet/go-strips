package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/theburntcrumpet/go-strip/src/config"
	"github.com/theburntcrumpet/go-strip/src/database"
	"github.com/theburntcrumpet/go-strip/src/routes"
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

	comicService, err := service.NewComicService(dbFactory, *conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	comicParser := &service.CbzComicParser{}
	router := gin.Default()
	routes.RegisterComicRoutes(router.Group("/api"), comicService, comicParser)
	router.Run()
}
