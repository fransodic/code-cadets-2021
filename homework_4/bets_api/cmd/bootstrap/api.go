package bootstrap

import (
	"code-cadets-2021/homework_4/bets_api/cmd/config"
	"code-cadets-2021/homework_4/bets_api/internal/api"
	"code-cadets-2021/homework_4/bets_api/internal/api/controllers"
	"code-cadets-2021/homework_4/bets_api/internal/api/controllers/validators"
	"code-cadets-2021/homework_4/bets_api/internal/domain/mappers"
	"code-cadets-2021/homework_4/bets_api/internal/domain/services"
	"code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite"
)

func newBetRequestValidator() *validators.BetRequestValidator {
	return validators.NewBetRequestValidator()
}

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, betMapper)
}

func newBetService(betRepository services.BetRepository) *services.BetService {
	return services.NewBetService(betRepository)
}

func newController(betRequestValidator controllers.BetRequestValidator, betService controllers.BetService, betMapper controllers.BetMapper) *controllers.Controller {
	return controllers.NewController(betRequestValidator, betService, betMapper)
}

// Api bootstraps the http server.
func Api(databaseExecutor sqlite.DatabaseExecutor) *api.WebServer {
	betRequestValidator := newBetRequestValidator()

	betMapper := newBetMapper()

	betRepository := newBetRepository(databaseExecutor, betMapper)
	betService := newBetService(betRepository)

	controller := newController(betRequestValidator, betService, betMapper)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
