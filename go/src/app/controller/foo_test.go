package controller

import (
	"app/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	is := assert.New(t)
	e := echo.New()
	c := Foo{}
	c.SetupFoo(e)
	ts := httptest.NewServer(e)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/api/v1/foo")
	is.NoError(err)
	v := model.Foo{}
	status := JsonParseResponse(res, &v)
	is.Equal(http.StatusOK, status)
	is.Equal("title", v.Title)
	is.Equal("name", v.Name)
}
