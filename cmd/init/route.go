package init

import (
	"ccvs/controller"
	"ccvs/services"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	g := router.Group("/api")
	c := controller.NewCreditCardControllers(services.NewCreditCardData(datastore))
	creditcardGroup := g.Group("/credit-card")

	// TODO : define Authentication in middleware package
	// uncomment this line for authentication
	// creditcardGroup.Use(middleware.Authentication())

	// handlers
	creditcardGroup.POST("/add", c.AddCreditCard)
	creditcardGroup.GET("/getall", c.GetCreditCards)
	creditcardGroup.GET("/get/:id", c.GetCreditCard)

}
