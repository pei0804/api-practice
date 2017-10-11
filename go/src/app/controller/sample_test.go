package controller

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	t.Run("echoHandler", func(t *testing.T) {
		candidates := []struct {
			url      string
			expected string
		}{
			{"http://localhost:8080/api/v1/sample", `{"title":"title","name":"name"}`},
			{"http://localhost:8080/api/v1/sample", `{"title":"title","name":"name"}`},
			{"http://localhost:8080/api/v1/sample", `{"title":"title","name":"name"}`},
		}
		for _, ca := range candidates {
			ca := ca
			t.Run(ca.url, func(t *testing.T) {
				t.Parallel()
				req := httptest.NewRequest(echo.GET, ca.url, nil)
				req.Header.Set(echo.HeaderAuthorization, "auth")
				rec := httptest.NewRecorder()
				e := echo.New()
				c := Sample{}
				ctx := e.NewContext(req, rec)
				if assert.NoError(t, c.Get(ctx)) {
					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, ca.expected, rec.Body.String())
				}
			})
		}
	})

}
