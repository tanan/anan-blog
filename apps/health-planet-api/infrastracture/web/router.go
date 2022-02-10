package web

import (
	"health-planet-api/config"
	"health-planet-api/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func LoadRouter(con *config.Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"*"},
	}))

	h, err := handler.NewHandler(con.ContoroDB.GetConnInfo())
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	e.GET("/systems/ping", ping)
	v1 := e.Group("v1")
	v1.GET("/me", func(ctx echo.Context) error { return h.Me(ctx) })
	v1.GET("/contracts/:id", func(ctx echo.Context) error { return h.GetContractById(ctx) })
	return e
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}