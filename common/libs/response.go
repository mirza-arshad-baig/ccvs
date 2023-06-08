package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ErrorResp struct {
	Error string `json:"error"`
}

// BuildResponse will build the json response for webservices
func BuildResponse(c *gin.Context, data interface{}, err error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		log.Errorf("Error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
	}
}
