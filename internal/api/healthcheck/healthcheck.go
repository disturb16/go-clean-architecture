package healthcheck

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/helper"
	logger "github.com/sanservices/apilogger/v2"
)

// Handler is the handler for healthchecks
type Handler struct {
}

// NewHandler is healthcheck Handler constructor
func NewHandler() *Handler {
	return &Handler{}
}

// RegisterRoutes initializes api routes
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.GET("/healthcheck", h.healthCheck)
}

func (h *Handler) healthCheck(c echo.Context) error {

	var err error

	resp := struct {
		Host     string    `json:"host"`
		Datetime time.Time `json:"datetime"`
	}{
		Datetime: time.Now(),
	}

	resp.Host, err = os.Hostname()
	if err != nil {
		logger.Error(c.Request().Context(), logger.LogCatHealth, err)
		return helper.RespondError(c, http.StatusOK, err)
	}

	return helper.RespondOk(c, resp)
}
