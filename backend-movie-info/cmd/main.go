package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var version string = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var configObj config

	flag.IntVar(&configObj.port, "port", 5000, "Server port to listen on")
	flag.StringVar(&configObj.env, "env", "development", "env development | production | staging")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: configObj,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", configObj.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	logger.Println("Starting the server on prot:", configObj.port)

	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}
