package app

import (
	"myapp/service"
	"myapp/util/logger"
)

type App struct {
	logger  logger.LoggerInterface
	svcBook service.BookServiceInterface
}

func NewApp(
	logger logger.LoggerInterface,
	svcBook service.BookServiceInterface,
) *App {
	return &App{
		logger:  logger,
		svcBook: svcBook,
	}
}

func (app *App) Logger() logger.LoggerInterface {
	return app.logger
}
