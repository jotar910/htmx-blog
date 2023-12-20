package handlers

import (
	"github.com/gin-gonic/gin"
	cerrors "github.com/jotar910/htmx-templ/pkg/errors"
	"github.com/jotar910/htmx-templ/pkg/logger"
)

func handleError(c *gin.Context, err error) {
	logger.L.Debugf("[ERROR]: %s", err)
	cerr := cerrors.Unwrap(err)
	c.JSON(cerr.Code, cerr)
}
