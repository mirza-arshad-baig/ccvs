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

// BuildResponse will build the json response for webservices
func BuildResponse(c *gin.Context, httpStatus int, data interface{}, err error) {
	if err == nil {
		c.JSON(httpStatus, SuccessResponse{Data: data, Code: httpStatus})
	} else {
		log.Errorf("Error: %s", err.Error())
		c.JSON(httpStatus, ErrorResp{Error: err.Error()})
	}
}
