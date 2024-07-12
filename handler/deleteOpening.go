package handler

import (
	"fmt"
	"github.com/febrandelli/gojobs/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		logger.Errorf("id is required")
		sendErrorResponse(c, http.StatusBadRequest, "id is required")
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Warningf("opening not found")
		sendErrorResponse(c, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening, "id = ?", id).Error; err != nil {
		logger.Errorf("error deleting opening: %v", err)
		sendErrorResponse(c, http.StatusInternalServerError, "error deleting opening")
		return
	}

	sendSuccessResponse(c, "delete-opening", opening)
}
