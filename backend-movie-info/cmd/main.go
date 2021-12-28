package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/rahulrana95/backend-movie-info/cmd/models"
)

var version string = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var configObj config

	flag.IntVar(&configObj.port, "port", 5000, "Server port to listen on")
	flag.StringVar(&configObj.env, "env", "development", "env development | production | staging")
	flag.StringVar(&configObj.db.dsn, "dsn", "postgres://rahulrana@localhost/movies-info-db?sslmode=disable", "Postgreql utl")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(configObj)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	app := &application{
		config: configObj,
		logger: logger,
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", configObj.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	logger.Println("Starting the server on prot:", configObj.port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}

func openDB(configObject config) (*sql.DB, error) {
	db, err := sql.Open("postgres", configObject.db.dsn)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}
