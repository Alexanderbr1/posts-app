package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"posts-app/internal/config"
	"posts-app/internal/repository"
	"posts-app/internal/server"
	"posts-app/internal/service"
	grpc_client "posts-app/internal/transport/grpc"
	"posts-app/internal/transport/rest"
	"posts-app/pkg/cache"
	"posts-app/pkg/database"
	"syscall"
)

func main() {
	cfg, err := config.New("configs", "main")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sqlx.DB) {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	}(db)

	c := cache.New()

	logsClient, err := grpc_client.NewClient(cfg.Grpc.Host, cfg.Grpc.Port)
	if err != nil {
		log.Fatal(err)
	}

	postsRepo := repository.NewRepository(db)
	postsService := service.NewService(cfg, c, postsRepo, logsClient)
	handler := rest.NewHandler(postsService)

	srv := server.New(cfg, handler.InitRoutes())
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal("error occurred while running http server: %s", err.Error())
		}
	}()

	log.Print("Post-app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Post-app stopped")

	if err := srv.Stop(context.Background()); err != nil {
		log.Fatal("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatal("error occurred on db connection close: %s", err.Error())
	}
}
