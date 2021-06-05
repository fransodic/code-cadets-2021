package controllers

import (
	"net/http"

	"code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betAcceptanceValidator BetAcceptanceValidator
	betAcceptanceService   BetAcceptanceService
}

// NewController creates a new instance of Controller
func NewController(betAcceptanceValidator BetAcceptanceValidator, betAcceptanceService BetAcceptanceService) *Controller {
	return &Controller{
		betAcceptanceValidator: betAcceptanceValidator,
		betAcceptanceService:   betAcceptanceService,
	}
}

// AcceptBet handles bet acceptance request.
func (e *Controller) AcceptBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betAcceptanceRequestDto models.BetAcceptanceRequestDto
		err := ctx.ShouldBindWith(&betAcceptanceRequestDto, binding.JSON)
		if err != nil {
			ctx.String(http.StatusBadRequest, "bet acceptance request is not valid.")
			return
		}

		if !e.betAcceptanceValidator.BetIsValid(betAcceptanceRequestDto) {
			ctx.String(http.StatusBadRequest, "bet acceptance request is not valid.")
			return
		}

		err = e.betAcceptanceService.PublishBet(betAcceptanceRequestDto.CustomerId, betAcceptanceRequestDto.SelectionId,
			betAcceptanceRequestDto.SelectionCoefficient, betAcceptanceRequestDto.Payment)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.Status(http.StatusOK)
	}
}
