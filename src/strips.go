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
	comic, err := service.CbzComicParser{}.ParseComic("bobby_make_believe_sample.cbz")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(comic.Filename)
	for _, page := range comic.Pages {
		fmt.Println(page.Filename)
		// print last 10 bytes of the page
		fmt.Println(page.Data[len(page.Data)-10:])
	}
}
