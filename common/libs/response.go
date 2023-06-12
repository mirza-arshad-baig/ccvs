package libs

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

// BuildResponse builds the JSON response for web services.
// It takes the Gin context, HTTP status code, response data, and error as input.
// If there is no error, it returns the success response with the data and HTTP status code.
// If there is an error, it logs the error and returns the error response with the error message and HTTP status code.
func BuildResponse(c *gin.Context, httpStatus int, data interface{}, err error) {
	if err == nil {
		c.JSON(httpStatus, SuccessResponse{Data: data, Code: httpStatus})
	} else {
		log.Errorf("Error: %s", err.Error())
		c.JSON(httpStatus, ErrorResp{Error: err.Error()})
	}
}
