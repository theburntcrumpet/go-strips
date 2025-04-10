package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Replace with your React app's URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.RegisterComicRoutes(router.Group("/api"), comicService, comicParser)
	router.Run()
}
