package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/zoshigayan/feedhell/db"
	"github.com/zoshigayan/feedhell/models"
	"github.com/zoshigayan/feedhell/services"
	"log"
	"net/http"
)

type FeedController struct {
}

func (fc FeedController) Init(g *echo.Group) {
	g.GET("", fc.Index)
	g.GET("/:id", fc.Show)
	g.GET("/new", fc.New)
	g.POST("/new", fc.Create)
}

func (fc FeedController) Index(c echo.Context) error {
	db := db.DbManager()

	feeds := []models.Feed{}
	db.Find(&feeds)

	return c.Render(http.StatusOK, "feed/index", map[string]interface{}{
		"Feeds": feeds,
	})
}

func (fc FeedController) Show(c echo.Context) error {
	db := db.DbManager()

	feed := models.Feed{}
	db.First(&feed, c.Param("id"))

	entries := []models.Entry{}
	db.Where("feed_id = ?", feed.ID).Find(&entries)

	return c.Render(http.StatusOK, "feed/show", map[string]interface{}{
		"Feed":    feed,
		"Entries": entries,
	})
}

func (fc FeedController) New(c echo.Context) error {
	return c.Render(http.StatusOK, "feed/new", map[string]interface{}{})
}

func (fc FeedController) Create(c echo.Context) error {
	db := db.DbManager()

	title, description, err := services.GetFeedInfo(c.FormValue("Link"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	feed := models.Feed{
		Title:       title,
		Description: description,
		Link:        c.FormValue("Link"),
	}
	result := db.Create(&feed)
	log.Println(result.Error)
	return c.Redirect(http.StatusSeeOther, "/feeds")
}
