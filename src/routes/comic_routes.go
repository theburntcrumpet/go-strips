package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theburntcrumpet/go-strip/src/controllers"
	"github.com/theburntcrumpet/go-strip/src/service"
)

func RegisterComicRoutes(router *gin.RouterGroup, comicService service.ComicService, comicParser service.ComicParser) {
	router.GET("/comics", func(c *gin.Context) {
		controllers.GetComics(c, comicService)
	})
	router.GET("/comics/:id/:page", func(c *gin.Context) {
		controllers.GetComicPage(c, comicService, comicParser)
	})
	router.GET("/comics/preview/:uuid", func(c *gin.Context) {
		controllers.GetPreviewImage(c, comicService)
	})
}
