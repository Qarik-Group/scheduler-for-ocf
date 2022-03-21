package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/scheduler-for-ocf/core"
	"github.com/starkandwayne/scheduler-for-ocf/http/routes"
)

func Server(bind string, services *core.Services) *http.Server {
	e := echo.New()

	// not technically necessary, just part of my default API skeleton
	e.GET("/health", func(c echo.Context) error {
		return c.String(
			http.StatusOK,
			"OK",
		)
	})

	routes.Apply(e, services)

	server := e.Server
	server.Addr = bind

	return server
}
