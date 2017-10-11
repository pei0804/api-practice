package controller

import (
	"app/model"
	"net/http"

	"app/mymiddleware"

	"github.com/labstack/echo"
)

type Sample struct {
}

func (c *Sample) SetupSample(e *echo.Echo) {
	auth := e.Group("/api/v1/sample", mymiddleware.Auth)
	auth.GET("", c.Get)
}

func (c *Sample) Get(ctx echo.Context) error {
	s := model.Sample{}
	sa := s.Get()
	return ctx.JSON(http.StatusOK, sa)
}
