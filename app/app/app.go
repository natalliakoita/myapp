package app

import (
	"myapp/service"
	"myapp/util/logger"
)

type App struct {
	logger  *logger.Logger
	svcBook service.BookServiceInterface
}

func NewApp(
	logger *logger.Logger,
	svcBook service.BookServiceInterface,
) *App {
	return &App{
		logger:  logger,
		svcBook: svcBook,
	}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}
