package order

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	order := &Order{}
	e.GET("/orders", order.Get)
	e.GET("/orders/:id", order.GetbyId)
	e.POST("/orders", order.Persist)
	e.PUT("/orders/:id", order.Put)
	e.DELETE("/orders/:id", order.Remove)
	return e
}
