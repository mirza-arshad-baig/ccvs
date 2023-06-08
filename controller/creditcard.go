package controller

import (
	"ccvs/common/libs"
	"ccvs/model"
	"ccvs/services"

	"github.com/gin-gonic/gin"
)

type CreditCardControllers struct {
	dataService services.ICreditCard
}

func NewCreditCardControllers(s services.ICreditCard) *CreditCardControllers {
	return &CreditCardControllers{dataService: s}
}

// AddCreditCard: AddCreditCard add credit card in DB
func (c *CreditCardControllers) AddCreditCard(ctx *gin.Context) {
	var (
		cdReq model.AddCreditCardReq
		err   error
	)
	err = ctx.BindJSON(&cdReq)
	if err == nil {
		err = c.dataService.AddCreditCard(ctx, cdReq)
	}
	libs.BuildResponse(ctx, nil, err)
}

// GetCreditCard: get credit card details by ID
func (c *CreditCardControllers) GetCreditCard(ctx *gin.Context) {
	creditCardID := ctx.Param("id")
	creditCardData, err := c.dataService.GetCreditCard(ctx, creditCardID)
	libs.BuildResponse(ctx, creditCardData, err)

}

// GetCreditCards: get all credit card details
func (c *CreditCardControllers) GetCreditCards(ctx *gin.Context) {
	creditCardData, err := c.dataService.GetCreditCards(ctx)
	libs.BuildResponse(ctx, creditCardData, err)
}
