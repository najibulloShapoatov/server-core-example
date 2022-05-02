package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/najibulloShapoatov/server-core-example/service"
	"github.com/najibulloShapoatov/server-core/monitoring/log"
	"github.com/najibulloShapoatov/server-core/server"
	"github.com/najibulloShapoatov/server-core/server/session"
	"github.com/najibulloShapoatov/server-core/settings"
	"github.com/najibulloShapoatov/server-core/utils"
)

var (
	configFile    = flag.String("config", "api.conf", "Configuration file")
	port          = flag.Int("port", 0, "Server port")
	logConfig     log.Config
	serverConfig  server.Config
	serviceConfig service.Config
	sessionConfig session.Config
)

func main() {
	flag.Parse()
	start()
}

func start() {
	defer fmt.Println("<<<BYE>>>")

	now := time.Now()

	config := settings.GetSettings()
	if err := config.Load(
		settings.NewFileLoader(*configFile, false),
		settings.NewEnvLoader(false, ""),
	); err != nil {
		fmt.Printf("Error loading settings: %s\n", err)
		os.Exit(1)
	}

	// Setting up logging
	if err := config.Unmarshal(&logConfig); err != nil {
		fmt.Printf("Error unmarshal log config: %s\n", err)
	}

	if err := log.Setup(logConfig); err != nil {
		fmt.Printf("Error setting up logging: %s\n", err)
	}

	// Setting up service
	if err := config.Unmarshal(&serviceConfig); err != nil {
		log.Errorf("Error unmarshal service config: %s", err)
	}
	apiService, err := service.New(&serviceConfig)
	if err != nil {
		log.Fatalf("Error initializing api service: %s", err)
	}

	// Setting up server
	if err := config.Unmarshal(&serverConfig); err != nil {
		log.Errorf("Error unmarshal server config: %s", err)
	}
	if port != nil && *port != 0 {
		serverConfig.Port = *port
	}
	svc, err := server.New(&serverConfig)
	if err != nil {
		log.Fatalf("Error creating HTTP server: %s", err)
	}

	if err = config.Unmarshal(&sessionConfig); err != nil {
		log.Errorf("Error unmarshal session config: %s", err)
	}

	// Register endpoints
	if err := server.RegisterRoute(apiService); err != nil {
		log.Fatalf("Error register service routes: %s", err)
	}

	// Start webserver
	if err := svc.Start(); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	log.Infof("Up and running (%s)", time.Since(now))

	// Stop webserver
	log.Infof("Got %s signal. Shutting down...", <-utils.WaitTermSignal())

	if err := svc.Stop(); err != nil {
		log.Errorf("Error stopping server: %s", err)
	}
}
