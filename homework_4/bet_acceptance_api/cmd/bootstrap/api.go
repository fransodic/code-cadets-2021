package bootstrap

import (
	"code-cadets-2021/homework_4/bet_acceptance_api/cmd/config"
	"code-cadets-2021/homework_4/bet_acceptance_api/internal/api"
	"code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers"
	"code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/validators"
	"code-cadets-2021/homework_4/bet_acceptance_api/internal/domain/services"
	"code-cadets-2021/homework_4/bet_acceptance_api/internal/infrastructure/rabbitmq"
	"github.com/streadway/amqp"
)

func newBetAcceptanceValidator() *validators.BetAcceptanceValidator {
	return validators.NewBetAcceptanceValidator(config.Cfg.BetValidator.CoefficientUpperBound, config.Cfg.BetValidator.PaymentLowerBound, config.Cfg.BetValidator.PaymentUpperBound)
}

func newBetReceivedPublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.BetReceivedPublisher {
	return rabbitmq.NewEventUpdatePublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherBetReceivedQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newBetAcceptanceService(publisher services.BetReceivedPublisher) *services.BetAcceptanceService {
	return services.NewBetAcceptanceService(publisher)
}

func newController(betAcceptanceValidator controllers.BetAcceptanceValidator, betAcceptanceService controllers.BetAcceptanceService) *controllers.Controller {
	return controllers.NewController(betAcceptanceValidator, betAcceptanceService)
}

// Api bootstraps the http server.
func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	betAcceptanceValidator := newBetAcceptanceValidator()
	betReceivedPublisher := newBetReceivedPublisher(rabbitMqChannel)
	betAcceptanceService := newBetAcceptanceService(betReceivedPublisher)
	controller := newController(betAcceptanceValidator, betAcceptanceService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
