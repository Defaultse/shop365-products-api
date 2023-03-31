// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"shop365-products-api/config"
	v1 "shop365-products-api/internal/controller/http/v1"
	"shop365-products-api/internal/usecase"
	"shop365-products-api/internal/usecase/adminuc"
	"shop365-products-api/internal/usecase/repo"
	"shop365-products-api/internal/usecase/repo/adminrepo"
	"shop365-products-api/internal/validator"
	"shop365-products-api/pkg/httpserver"
	"shop365-products-api/pkg/logger"
	"shop365-products-api/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	v := validator.NewValidator()

	a := cfg.PG
	fmt.Println(a)
	// Connecting to postgres
	pgMaps, err := postgres.NewPostgres(context.TODO(), cfg.PG)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - Connecting to postgres: %w", err))
	}

	// sqlDB, err := pgClient.DB()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// sqlDB.SetMaxOpenConns(5)

	// Connecting to mongo
	mongoCreds := options.Credential{Username: cfg.MONGO.Username, Password: cfg.MONGO.Password}
	ctx := context.Background()
	mongoCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(cfg.MONGO.URL).SetAuth(mongoCreds))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - Connecting to mongo: %w", err))
		return
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("err", err)
		return
	}

	// Use case
	allUseCase := &usecase.AllUseCases{
		CategoryUC: *usecase.NewCategoryUC(
			repo.NewCategoryRepo(mongoClient),
		),
		ProductUC: *usecase.NewProductUC(
			repo.NewProductRepo(pgMaps),
		),
		AdminProductUC: *adminuc.NewProductUC(
			adminrepo.NewAdminProductRepo(pgMaps),
		),
	}

	// RabbitMQ RPC Server
	// rmqRouter := amqprpc.NewRouter(translationUseCase)

	// rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	// }

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, v, *allUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		// case err = <-rmqServer.Notify():
		// 	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	// err = rmqServer.Shutdown()
	// if err != nil {
	// 	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	// }
}
