package controller

import (
	"ccvs/common/libs"
	"ccvs/model"
	"ccvs/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreditCardControllers struct {
	dataService services.ICreditCard
}

func NewCreditCardControllers(s services.ICreditCard) *CreditCardControllers {
	return &CreditCardControllers{dataService: s}
}

// AddCreditCard adds a credit card to the database.
func (c *CreditCardControllers) AddCreditCard(ctx *gin.Context) {
	var (
		cdReq model.AddCreditCardReq
		err   error
	)

	if err = ctx.BindJSON(&cdReq); err != nil {
		libs.BuildResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	if err = c.dataService.AddCreditCard(ctx, cdReq); err != nil {
		libs.BuildResponse(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	libs.BuildResponse(ctx, http.StatusOK, "Credit card added successfully", err)
}

// GetCreditCard retrieves the credit card details by ID.
func (c *CreditCardControllers) GetCreditCard(ctx *gin.Context) {
	creditCardID := ctx.Param("id")

	creditCardData, err := c.dataService.GetCreditCard(ctx, creditCardID)
	if err != nil {
		libs.BuildResponse(ctx, http.StatusInternalServerError, nil, err)
		return
	}
	libs.BuildResponse(ctx, http.StatusOK, creditCardData, err)
}

// GetCreditCards retrieves all credit card details.
func (c *CreditCardControllers) GetCreditCards(ctx *gin.Context) {
	creditCardData, err := c.dataService.GetCreditCards(ctx)
	if err != nil {
		libs.BuildResponse(ctx, http.StatusInternalServerError, nil, err)
		return
	}

	libs.BuildResponse(ctx, http.StatusOK, creditCardData, err)
}
