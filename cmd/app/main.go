package main

import (
	"fmt"
	"myapp/app/router"
	"myapp/config"
	"myapp/repository"
	"myapp/service"
	lr "myapp/util/logger"
	"net/http"

	dbConn "myapp/adapter/gorm"
	"myapp/app/app"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	conn, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}
	if appConf.Debug {
		conn.LogMode(true)
	}

	db := repository.NewBookRepo(conn)
	svcBook := service.NewBookService(db)

	// validator := validator.New()

	application := app.NewApp(logger, svcBook)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
