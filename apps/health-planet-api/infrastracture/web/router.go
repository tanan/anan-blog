package web

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"health-planet-api/config"
	"health-planet-api/handler"
	"net/http"
)

var han *handler.Handler

func LoadRouter(con *config.Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"*"},
	}))

	han, _ = handler.NewHandler(con, NewClient(5))

	e.GET("/systems/ping", ping)
	v1 := e.Group("v1")
	v1.Use(auth)
	v1.GET("/me", func(ctx echo.Context) error { return han.Me(ctx) })
	return e
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, refreshToken, err := han.GetAuth()
		if err != nil {
			c.Error(err)
		}
		han.APIConfig.AccessToken = token
		han.APIConfig.RefreshToken = refreshToken
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
