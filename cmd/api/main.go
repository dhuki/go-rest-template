package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	httpHandler "github.com/dhuki/go-rest-template/internal/adapter/http"
	v1 "github.com/dhuki/go-rest-template/internal/adapter/http/v1"
	"github.com/dhuki/go-rest-template/internal/infra/configloader"
	postgres "github.com/dhuki/go-rest-template/internal/infra/database/postgres"
	"github.com/dhuki/go-rest-template/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	env := os.Getenv("ENV")
	if env == "" {
		env = "LOCAL"
	}
	flag.StringVar(&configloader.Conf.App.Env, "env", env, "define environment")
	flag.Parse()

	// init config
	configloader.InitConsul(ctx, configloader.Conf.App.Env)

	// init database
	pgRepository, err := postgres.NewPostgreSQLClient().NewPgRepository(&configloader.Conf.ConnDatabase)
	if err != nil {
		logger.Fatal(ctx, "postgre.ConnectDatabase", "Error connect to database postgre, err : %v", err)
	}

	// if mysqlDB, err := mysql.NeMySQLClient().ConnectDatabase(&configloader.Conf.ConnDatabase); err != nil {
	// 	logger.Fatal(ctx, "mysql.ConnectDatabase", "Error connect to database mysql, err : %v", err)
	// }

	// init dependency
	handler := httpHandler.NewHandler(pgRepository)

	// init router
	v1Server := v1.NewHttpHandlerV1(handler, configloader.Conf.App.Port)

	// run server v1
	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		logrus.Infof("We received an interrupt signal, shutting down service")
		if err := v1Server.Stop(ctx); err != nil {
			logger.Fatal(ctx, "route stop", "Error stopping service go-date, err : %v", err)
		}
		logger.Info(ctx, "route stop", "Success stopping service go-rest-example")
		close(idleConsClosed)
	}()

	logger.Info(ctx, "route start", "Success start service go-rest-example listening on port :%d", configloader.Conf.App.Port)
	if err := v1Server.Start(ctx); err != nil {
		logger.Fatal(ctx, "route start", "Error starting service go-rest-example, err : %v", err)
	}
	<-idleConsClosed

}
