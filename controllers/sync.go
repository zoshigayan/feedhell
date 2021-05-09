package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/zoshigayan/rss_reader/services"
	"net/http"
)

type SyncController struct{}

func (sc SyncController) Init(g *echo.Group) {
	g.POST("", sc.Sync)
}

func (sc SyncController) Sync(c echo.Context) error {
	err := services.Sync()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}
