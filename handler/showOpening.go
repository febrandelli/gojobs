package handler

import (
	"fmt"
	"github.com/febrandelli/gojobs/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowOpeningHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		logger.Errorf("id is required")
		sendErrorResponse(c, http.StatusBadRequest, "id is required")
		return
	}

	var opening schema.Opening

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("opening id: %v notfound", err)
		sendErrorResponse(c, http.StatusNotFound, fmt.Sprintf("opening id: %s not found", id))
		return
	}

	sendSuccessResponse(c, "show-opening", opening)
}
