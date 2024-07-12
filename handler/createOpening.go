package handler

import (
	"github.com/febrandelli/gojobs/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}

	if err := c.BindJSON(&request); err != nil {
		logger.Errorf("mapping request body error: %v", err.Error())
		sendErrorResponse(c, http.StatusBadRequest, "mapping request body error")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendErrorResponse(c, http.StatusBadRequest, "validation error")
		return
	}

	opening := schema.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("creating opening error: %v", err.Error())
		sendErrorResponse(c, http.StatusInternalServerError, "error create opening on database")
		return
	}

	sendSuccessResponse(c, "create-opening", opening)
}
