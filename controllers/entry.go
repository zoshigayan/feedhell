package controllers

import (
	"github.com/labstack/echo/v4"
	// "github.com/zoshigayan/rss_reader/db"
	// "github.com/zoshigayan/rss_reader/models"
	"net/http"
)

type EntryController struct {
}

func (ec EntryController) Init(g *echo.Group) {
	g.GET("/", ec.Index)
}

func (fc EntryController) Index(c echo.Context) error {
	return c.String(http.StatusOK, "entries")
	// db := db.DbManager()

	// entries := []models.Entry{}
	// db.Find(&entries)

	// return c.Render(http.StatusOK, "entry/index", map[string]interface{}{
	// 	"Entries": entries,
	// })
}
