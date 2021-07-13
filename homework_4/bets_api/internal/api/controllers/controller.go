package controllers

import (
	"code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct {
	BetRequestValidator BetRequestValidator
	betService          BetService
	betMapper           BetMapper
}

func NewController(betRequestValidator BetRequestValidator, betService BetService, betMapper BetMapper) *Controller {
	return &Controller{
		BetRequestValidator: betRequestValidator,
		betService:          betService,
		betMapper:           betMapper,
	}
}

func (c *Controller) GetBetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if !isValidID(id) {
			log.Println("Invalid id")
			ctx.Status(http.StatusBadRequest)
			return
		}

		resultBet, exists, err := c.betService.GetByID(ctx, id)
		if err != nil {
			log.Println("Failed to fetch a bet, error: ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if !exists {
			log.Println("No bets with id: ", id)
			ctx.Status(http.StatusNotFound)
		} else {
			betDto := c.betMapper.MapResultToDto(resultBet)
			ctx.JSON(http.StatusOK, betDto)
		}
	}
}

func isValidID(id string) bool {
	return id != ""
}

func (c *Controller) GetBetsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Query("status")
		if !c.BetRequestValidator.StatusIsValid(status) {
			ctx.Status(http.StatusBadRequest)
			return
		}

		resultBets, err := c.betService.GetBetsByStatus(ctx, status)
		if err != nil {
			log.Println("Failed to fetch a bet, error: ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if len(resultBets) == 0 {
			log.Println("No bets with status: ", status)
			ctx.Status(http.StatusNotFound)
		} else {
			betDto := c.mapResultsToDto(resultBets)
			ctx.JSON(http.StatusOK, betDto)
		}
	}
}

func (c *Controller) GetBetsByCustomerID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerId := ctx.Param("id")
		if !isValidID(customerId) {
			ctx.Status(http.StatusBadRequest)
			return
		}

		resultBets, err := c.betService.GetBetsByCustomerID(ctx, customerId)
		if err != nil {
			log.Println("Failed to fetch a bet, error: ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if len(resultBets) == 0 {
			log.Println("No bets for user id: ", customerId)
			ctx.Status(http.StatusNotFound)
		} else {
			betDto := c.mapResultsToDto(resultBets)
			ctx.JSON(http.StatusOK, betDto)
		}
	}
}

func (c *Controller) mapResultsToDto(bets []domainmodels.Bet) []models.BetResponseDto {
	var betDto []models.BetResponseDto
	for _, bet := range bets {
		betDto = append(betDto, c.betMapper.MapResultToDto(bet))
	}

	return betDto
}
