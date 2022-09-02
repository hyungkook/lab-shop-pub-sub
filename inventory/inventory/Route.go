package inventory

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	inventory := &Inventory{}
	e.GET("/inventories", inventory.Get)
	e.GET("/inventories/:id", inventory.GetbyId)
	e.POST("/inventories", inventory.Persist)
	e.PUT("/inventories/:id", inventory.Put)
	e.DELETE("/inventories/:id", inventory.Remove)
	return e
}
