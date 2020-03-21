package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MrJootta/GoUrl/handler"
	"github.com/MrJootta/GoUrl/internal/config"
	"github.com/MrJootta/GoUrl/internal/storage/mysql"
)

func main() {
	conf := config.StartConfigs()

	db, err := mysql.New(conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPass, conf.DBDatabase)
	if err != nil {
		log.Fatal("database connection error")
	}
	defer db.Close()

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", conf.ServerHost, conf.ServerPort),
		Handler:      handler.New(db),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	go func() {
		log.Printf("server start http://%s", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("error: %v", err)
		} else {
			log.Println("sever down")
		}
	}()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)

	sig := <-sigquit
	log.Printf("caught sig: %+v", sig)
	log.Printf("shutting down server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("Not able to shut down: %v", err)
	} else {
		log.Println("server closed")
	}
}
