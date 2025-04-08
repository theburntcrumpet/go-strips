package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theburntcrumpet/go-strip/src/service"
)

func GetComics(c *gin.Context, comicService service.ComicService) {
	// Get the list of comics from the service
	query := c.Query("query")
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	if query == "" && page == "" && pageSize == "" {
		comics, err := comicService.GetComics()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comics"})
			return
		}
		c.JSON(http.StatusOK, comics)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}
	comics, err := comicService.GetComicsPaginated(query, pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comics"})
		return
	}
	c.JSON(http.StatusOK, comics)
}

func GetComicPage(c *gin.Context, comicService service.ComicService, comicParser service.ComicParser) {
	// Get the comic ID from the URL parameter
	id := c.Param("id")
	page := c.Param("page")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comic ID is required"})
		return
	}

	if page == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page number is required"})
		return
	}

	// Convert the page number to an integer
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	// Get the comic from the service
	comic, err := comicService.GetComicWithId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comic"})
		return
	}

	// Parse the comic using the parser
	comicData, err := comicParser.ParseComic(comic.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse comic"})
		return
	}
	// Check if the page number is valid
	if pageInt < 0 || pageInt >= len(comicData.Pages) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", comicData.Pages[pageInt].Data)
}

func GetPreviewImage(c *gin.Context, comicService service.ComicService) {
	// Get the UUID from the URL parameter
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID is required"})
		return
	}

	// Get the preview image from the service
	image, err := comicService.GetPreviewImageWithUuid(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get preview image"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", image)
}
