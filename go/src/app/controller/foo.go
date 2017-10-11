package controller

import (
	"app/model"
	"net/http"

	"github.com/labstack/echo"
)

type Foo struct {
}

func (c *Foo) SetupFoo(e *echo.Echo) {
	e.GET("/api/v1/foo", c.Get)
}

func (c *Foo) Get(ctx echo.Context) error {
	s := model.Foo{}
	sa := s.Get()
	return ctx.JSON(http.StatusOK, sa)
}
