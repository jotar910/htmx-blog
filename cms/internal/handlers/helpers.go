package handlers

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	components_core "github.com/jotar910/buzzer-cms/internal/components/core"
	cerrors "github.com/jotar910/buzzer-cms/pkg/errors"
	"github.com/jotar910/buzzer-cms/pkg/logger"
	"net/http"
)

func render(c *gin.Context, html templ.Component) {
	if c.GetHeader("HX-Request") == "" {
		// This means it's the initial full page load
		// Run your specific middleware logic here
		// For example, initializing session data, etc.

		// Log for demonstration purposes
		c.HTML(http.StatusOK, "", components_core.Index(html))
	} else {
		c.HTML(http.StatusOK, "", html)
	}
}

func handleError(c *gin.Context, err error) {
	logger.L.Debugf("[ERROR]: %s", err)
	cerr := cerrors.Unwrap(err)
	c.JSON(cerr.Code, cerr)
}
