package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tags/internal/api"
	"tags/internal/storage"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	godotenv.Load()

	db, err := storage.NewPostgreStore()
	if err != nil {
		logrus.Fatal(err)
	}

	if err := db.Init(); err != nil {
		logrus.Fatal(err)
	}

	server := api.NewAPIServer(os.Getenv("PORT"), db)

	go func() {
		if err := server.Run(); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err)
		}
	}()

	logrus.Print("Articles app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Articles Shutdown Down")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutdown: %s", err.Error())
	}
}
