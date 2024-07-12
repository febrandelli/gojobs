package handler

import (
	"github.com/febrandelli/gojobs/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningsHandler(c *gin.Context) {
	var openings []schema.Opening
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error finding openings: %v", err)
		sendErrorResponse(c, http.StatusInternalServerError, "error finding openings")
	}

	sendSuccessResponse(c, "list-openings", openings)
}
