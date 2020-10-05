package internal

import (
	"net/http"

	"github.com/labstack/echo"
)

// Health checks health of the API
func (h *Handler) Health(c echo.Context) (err error) {
	return c.String(http.StatusOK, "OK")
}
